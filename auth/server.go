package main

// # auth - server for athentication and user-registration.
//
// ## The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT       - the server is listening on this portnumber
//   * DB_HOST           - database host for connecting the auth database
//   * DB_NAME           - database name for connecting the auth database
//   * DB_USER           - database user for connecting the auth database
//   * DB_PASSWORD       - database user-password for connecting the auth database
//   * P_FACEBOOK_KEY    - Facebook-key for login-provider (goth)
//   * P_FACEBOOK_SECRET - Facebook-secret for login-provider (goth)
//   * P_GPLUS_KEY    	 - Google+-key for login-provider (goth)
//   * P_GPLUS_SECRET    - Google+-secret for login-provider (goth)
//   * P_APPLE_KEY    	 - Apple-key for login-provider (goth)
//   * P_APPLE_SECRET    - Apple-secret for login-provider (goth)
//   * ACTIVATION_URL    - The url that holds the link to the activaton route. <- depends on dns and albert config for auth
//   * ACTIVATION_SUCCESS_URL - Thes urls that is called when activation has happened. Its a static page that is reached through redirect
//   * BASE_URL          - Defines the base URL for contructing i.e. callback-URLs. Should be the name of the server the app is running on.
//   * SESSION_SECRET    - Defines the secret that is uses to encrypt the sessions.
import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/damonkeys/goth/gothic"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/apple"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"

	"github.com/damonkeys/ch3ck1n/auth/models"

	"github.com/damonkeys/ch3ck1n/auth/applesupport"
	bubblesclient "github.com/damonkeys/ch3ck1n/monkeys/bubbles"
	"github.com/damonkeys/ch3ck1n/monkeys/config"
	"github.com/damonkeys/ch3ck1n/monkeys/database"
	l "github.com/damonkeys/ch3ck1n/monkeys/logger"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// ServerConfigStruct holds the server-config for auth
	ServerConfigStruct struct {
		Port          string                `env:"SERVER_PORT"`
		Baseurl       string                `env:"BASE_URL"`
		SessionSecret string                `env:"SESSION_SECRET"`
		Providers     ProvidersStruct       `json:"providers"`
		Database      database.ConfigStruct `json:"database"`
		Activation    ActivationStruct      `json:"activation"`
	}

	// ProvidersStruct defines all configured OAuth-providers
	ProvidersStruct struct {
		Facebook FacebookSecretsStruct `json:"facebook"`
		Gplus    GplusSecretsStruct    `json:"gplus"`
		Apple    AppleSecretsStruct    `json:"apple"`
	}
	// FacebookSecretsStruct defines one key-secret-pair for facebooks oauth-provider
	FacebookSecretsStruct struct {
		Key    string `env:"P_FACEBOOK_KEY" json:"key"`
		Secret string `env:"P_FACEBOOK_SECRET" json:"secret"`
	}
	// GplusSecretsStruct defines one key-secret-pair for Google+ oauth-provider
	GplusSecretsStruct struct {
		Key    string `env:"P_GPLUS_KEY" json:"key"`
		Secret string `env:"P_GPLUS_SECRET" json:"secret"`
	}
	// AppleSecretsStruct defines one key-secret-pair for Apple oauth-provider
	AppleSecretsStruct struct {
		Key    string `env:"P_APPLE_KEY" json:"key"`
		Secret string `env:"P_APPLE_SECRET" json:"secret"`
	}

	// LogoutUserResponse returns successfully logouts
	LogoutUserResponse struct {
		Successful bool `json:"successful"`
	}

	// ActivationStruct contains information about activation configuration.
	// Currently its the url where an account can be activated and
	// the url to which the user is redirected when activation was successful or is ongoing
	ActivationStruct struct {
		URL                string `env:"ACTIVATION_URL" json:"url"`
		AcitvationStateURL string `env:"ACTIVATION_STATE_URL" json:"activation_state_url"`
	}

	// LoginStatusResponse returns user login status
	LoginStatusResponse struct {
		UserOnline bool   `json:"useronline"`
		Username   string `json:"username"`
		Merchant   bool   `json:"merchant"`
		AvatarURL  string `json:"avatarurl"`
	}

	// UserInfo returns all userdata for profile
	UserInfo struct {
		PreferredLanguage string `json:"preferredLanguage"`
		UserOnline        bool   `json:"useronline"`
		Username          string `json:"username"`
		Merchant          bool   `json:"merchant"`
		Email             string `json:"email"`
		UUID              string `json:"uuid"`
		ActivationToken   string `json:"activationToken"`
	}

	// IntUserInfoResponse returns all userdata for internal server connections
	IntUserInfoResponse struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Successful bool   `json:"successful"`
	}

	// MerchantModeResponse returns success-message for setting merchant flag for user
	MerchantModeResponse struct {
		Successful bool `json:"successful"`
	}

	// AccountActivation carries the token and provides functions to activate a user/ trigger follow up actions
	AccountActivation struct {
		Token string `json:"token"`
	}

	// SessionData lets you access relevant session Data
	SessionData interface {
		findSessionUser(c echo.Context) (*models.User, error)
	}

	// SessionDataStruct implements SessionData
	SessionDataStruct struct {
	}

	// ActivationTokenGenerator generates an activationToken - currently only an UUID and stores it into the users data (DB)
	ActivationTokenGenerator interface {
		createAndStoreActivationTokenForUser(c echo.Context) error
	}

	// Activator activates an Account
	Activator interface {
		activate(c echo.Context) error
	}

	// LanguageConfigurator sets the language for something (most of the time this will be the current user)
	LanguageConfigurator interface {
		Savelanguage(c echo.Context)
	}
)

// SaveLanguage sets the language to the given string. Should be a string explaining the language in context of this system, e.g. 'en' for english.
func (userInfo *UserInfo) SaveLanguage(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()
	//parse params default sensible use LanguageHelper
	language := c.Param("l")
	tracing.LogString(span, fmt.Sprintf("Selected language for user %s", userInfo.UUID), language)
	sessionUser, err := models.FindUserByUUID(tracing.GetContext(c), userInfo.UUID)
	if err != nil {
		tracing.LogError(span, err)
		return err
	}
	sessionUser.PreferredLanguage = language
	db := database.DB.Save(sessionUser)
	if db.Error != nil {
		tracing.LogError(span, db.Error)
		return db.Error
	}
	return nil
}

// activate activates the account and redirects to login page
func (activation *AccountActivation) activate(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()
	foundUser := &models.User{}
	db := database.DB.Where("activation_token = ?", activation.Token).First(foundUser)
	if db.Error != nil {
		tracing.LogString(span, "could find user with this activation code", activation.Token)
		return db.Error
	}
	foundUser.Active = true
	foundUser.ActiveSince = time.Now()
	db = database.DB.Save(foundUser)
	if db.Error != nil {
		tracing.LogStruct(span, "could not update db user during activation", foundUser)
		return db.Error
	}
	return nil
}

func (userInfo *UserInfo) createAndStoreActivationTokenForUser(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()
	user := models.User{}
	db := database.DB.Where("email = ?", userInfo.Email).First(&user)
	if db.Error != nil {
		tracing.LogStruct(span, "could find user for activation for this data", userInfo)
		return db.Error
	}
	user.ActivationToken = uuid.New().String()
	user.ActivationTokenCreation = time.Now()
	db = database.DB.Save(user)
	if db.Error != nil {
		tracing.LogStruct(span, "could not update db user with activationToken", user)
		return db.Error
	}
	userInfo.PreferredLanguage = user.PreferredLanguage
	userInfo.ActivationToken = user.ActivationToken
	return nil

}

const sessionName = "_ch3ck1n_callback"

// ServerConfig defines the configuration for auth
var serverConfig ServerConfigStruct

func main() {
	// tracer init
	closer, span, ctx := tracing.InitJaeger("auth")
	defer closer.Close()

	// Init echo
	e := echo.New()
	l.ConfigureLogger(ctx, "auth", e)
	readEnvironmentConfig(ctx, e.Logger)
	tracing.LogStruct(span, "serverConfig", serverConfig)

	// Init goth-providers that we use
	e.Logger.Debug("Initialise Goth with Facebook, Google+ and Apple providers")
	tracing.LogString(span, "goth", "Initialise Goth with Facebook, Google+ and Apple providers")
	initGoth()

	if err := database.InitDatabase(serverConfig.Database); err != nil {
		e.Logger.Fatal(err)
		tracing.LogError(span, err)
		span.Finish()
		os.Exit(0)
	}
	defer database.DB.Close()

	// creeate session store for echo and gorilla (used by goth!)
	sessionStore := sessions.NewCookieStore([]byte(serverConfig.SessionSecret))
	gothic.Store = sessionStore
	e.Use(session.Middleware(sessionStore))
	e.Use(tracing.Middleware("auth"))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/login", login)
	e.GET("/logout", logout)
	callbackMethods := []string{"GET", "POST"}
	e.Match(callbackMethods, "/callback", callback)
	e.GET("/status", getLoginStatus)
	e.GET("/userInfos", getUserInfos)
	//echo does not handle different names for parameters for the same route but different http verbs correctly. Thats why we use ':param' everywhere
	e.POST("/activation/:param", sendActivationMail)
	e.GET("/activation/:param", processActivationState)
	e.POST("/merchant/activate", activateMerchant)
	e.POST("/merchant/deactivate", deactivateMerchant)
	e.POST("/user/:l", setPreferredLang)
	span.Finish()
	e.Logger.Fatal(e.Start(":" + serverConfig.Port))
}

// readEnvironmentConfig reads all needed environment variables and save it in ServerConfig struct
func readEnvironmentConfig(ctx context.Context, log echo.Logger) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	// read config from environment variables to struct
	configInterface, err := config.ReadEnvVars(ctx, ServerConfigStruct{})
	if err != nil {
		log.Error(err)
		tracing.LogError(span, err)
		os.Exit(-1)
	}
	serverConfig = configInterface.(ServerConfigStruct)
}

func initGoth() {
	goth.UseProviders(
		facebook.New(serverConfig.Providers.Facebook.Key, serverConfig.Providers.Facebook.Secret, "https://dev.checkin.chckr.de/auth/callback?provider=facebook"),
		gplus.New(serverConfig.Providers.Gplus.Key, serverConfig.Providers.Gplus.Secret, "https://dev.checkin.chckr.de/auth/callback?provider=gplus"),
		apple.New(serverConfig.Providers.Apple.Key, serverConfig.Providers.Apple.Secret, "https://dev.checkin.chckr.de/auth/callback", nil, apple.ScopeName, apple.ScopeEmail),
	)
}

// login calls given provider oauth and show the login dialog
func login(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	sess, _ := session.Get(sessionName, c)
	sess.Options = &sessions.Options{
		Path:     "/auth",
		MaxAge:   30,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	sessions.NewCookie("callbackURL", c.QueryParam("callbackUrl"), sess.Options)
	sess.Values["callbackURL"] = c.QueryParam("callbackUrl")
	sess.Save(c.Request(), c.Response())

	c.Logger().Debug("Cookies created.")
	tracing.LogString(span, "log", "Cookies created.")

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request()); err == nil {
		c.Logger().Debug("User already logged in.")
		tracing.LogString(span, "log", "User already logged in.")
		tracing.LogStruct(span, "gothUser", gothUser)
		tracing.LogString(span, "email", gothUser.Email)
		tracing.LogString(span, "name", gothUser.FirstName+" "+gothUser.LastName)
		tracing.LogString(span, "nickname", gothUser.NickName)
		tracing.LogString(span, "userid", gothUser.UserID)
		// user already logged in
		return createNewSessionCookie(c, gothUser)
	}
	// user has to login
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil //c.Redirect(http.StatusTemporaryRedirect, "https://dev.checkin.chckr.de")
}

// logout remove session cookie and sign out the user
func logout(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	logoutUserResponse := &LogoutUserResponse{
		Successful: false,
	}
	err := removeSessionCookie(c)
	if err != nil {
		c.Logger().Debugf("error while removing session-cookie: %s", err)
		tracing.LogError(span, err)
		return c.JSON(http.StatusBadRequest, logoutUserResponse)
	}
	logoutUserResponse.Successful = true
	return c.JSON(http.StatusOK, logoutUserResponse)
}

func resolveUserNameFromRequestIfApple(c echo.Context, user *goth.User) error {
	span := tracing.Enter(c)
	defer span.Finish()
	if user.Provider == "apple" {
		r := c.Request()
		userData := &applesupport.UserNameData{}
		if isExpectedPostRequest(r.URL.Query(), r.Method) {
			err := userData.ResolveUserNames(r)
			if err != nil {
				tracing.LogError(span, err)
				return err
			}
			if userData.FirstName != "" {
				user.FirstName = userData.FirstName
			}
			if userData.LastName != "" {
				user.LastName = userData.LastName
			}
			if userData.FirstName != "" && userData.LastName != "" {
				user.Name = userData.FirstName + " " + userData.LastName
			}
			return nil
		}
		tracing.LogStruct(span, "No apple user data", r)
		return nil
	}
	tracing.LogStruct(span, "No apple provider. Instead: provider:", user.Provider)
	return nil
}

// callback after successful oauth login
func callback(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		tracing.LogError(span, err)
	}
	err = resolveUserNameFromRequestIfApple(c, &user)
	if err != nil {
		tracing.LogError(span, err)
	}
	c.Logger().Debugf("callback-user-data: nickname [%s] / name [%s] / email [%s]", user.NickName, user.FirstName+" "+user.LastName, user.Email)
	tracing.LogStruct(span, "callback-user-data", user)
	if user.Email == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/NoEmail")
	}
	if err != nil {
		// there is a problem in login the user
		return c.Redirect(http.StatusTemporaryRedirect, getCallbackURLFromSession(c))
	}
	// user successfully logged in - create a JWT Token
	return createNewSessionCookie(c, user)
}

// getLoginStatus will be called by react-app to get the login status
func getLoginStatus(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	loginStatusResponse := &LoginStatusResponse{
		UserOnline: false,
	}
	sessionUser, err := findSessionUser(c)
	if err != nil {
		c.Logger().Info(err)
		tracing.LogError(span, err)
		return c.JSON(http.StatusOK, loginStatusResponse)
	}
	loginStatusResponse.UserOnline = true
	loginStatusResponse.Username = sessionUser.Name
	loginStatusResponse.AvatarURL = sessionUser.AvatarURL
	loginStatusResponse.Merchant = sessionUser.Merchant

	return c.JSON(http.StatusOK, loginStatusResponse)
}

func getSessionUserInfo(c echo.Context) (*UserInfo, error) {
	span := tracing.Enter(c)
	defer span.Finish()
	userInfo := &UserInfo{
		UserOnline: false,
	}
	sessionUser, err := findSessionUser(c)
	if err != nil {
		c.Logger().Info(err)
		tracing.LogError(span, err)
		return nil, err
	}
	userInfo.UserOnline = true
	userInfo.Username = sessionUser.Name
	userInfo.Email = sessionUser.Email
	userInfo.UUID = sessionUser.UUID
	userInfo.Merchant = sessionUser.Merchant

	return userInfo, nil
}

// getUserInfos returns all (session based) Userdata as JSON only if the user is online
func getUserInfos(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	userInfo, err := getSessionUserInfo(c)
	if err != nil {
		c.Logger().Info(err)
		tracing.LogError(span, err)
		return c.JSON(http.StatusOK, userInfo)
	}

	return c.JSON(http.StatusOK, userInfo)
}

func setPreferredLang(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	sessionUser, err := getSessionUserInfo(c)
	if err != nil {
		c.Logger().Info(err)
		tracing.LogError(span, err)
	}
	sessionUser.SaveLanguage(c)

	return c.JSON(http.StatusOK, "OK")
}

// sendActivationMail sends an activationMail for the given User (Session)
func sendActivationMail(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	email := c.Param("param")
	if email == "" {
		return c.JSON(http.StatusBadRequest, "missing data")
	}
	userInfo := UserInfo{
		Email: email,
	}

	span.SetTag("email", email)
	c.Logger().Debugf("Trying to create ActivationToken for user with email %s", email)

	err := userInfo.createAndStoreActivationTokenForUser(c)
	if err != nil {
		// Needs good error Page with Support or even better: telemetry alerts
		tracing.LogStruct(span, "activation cannot be created", userInfo)
		return c.JSON(http.StatusInternalServerError, "error")
	}

	preparer := prepareI18nActivationMessageText(userInfo)
	tracing.LogStruct(span, "activation-mail-data", preparer)
	sendMail(preparer)
	return c.JSON(http.StatusOK, userInfo.Email)
}

func prepareI18nActivationMessageText(userInfo UserInfo) MGMessagePreparer {
	if userInfo.PreferredLanguage == "de" {
		preparer := &CTAMailContext{
			templatename: "cta-tpl",
			recipient:    "support@chckr.de",
			sender:       "support@chckr.de",
			subject:      fmt.Sprintf("Zugangs-Aktivierung für %s", userInfo.Email),
			cta:          "Zugangs-Aktivierung notwendig",
			body:         "Du hast es fast geschafft: Klicke auf den Aktivierungslink um die Einrichtung Deines Zugangs abzuschließen. Vielen Dank <3",
			ctalink:      fmt.Sprintf(serverConfig.Activation.URL, userInfo.ActivationToken),
			linktext:     "Jetzt klicken und aktivieren",
		}
		return preparer
	}
	preparer := &CTAMailContext{
		templatename: "cta-tpl",
		recipient:    "support@chckr.de",
		sender:       "support@chckr.de",
		subject:      fmt.Sprintf("Account-Activation for %s", userInfo.Email),
		cta:          "Account-Activation necessary",
		body:         "You almost made it: Click on the activation link to complete the configuration of your account. Thank you <3",
		ctalink:      fmt.Sprintf(serverConfig.Activation.URL, userInfo.ActivationToken),
		linktext:     "Click now to activate",
	}
	return preparer

}

func processActivationState(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	span.SetTag("account-activation", "Traces the handling of the activation")
	activationtoken := c.Param("param")
	if activationtoken == "" {
		tracing.LogString(span, "activationtoken", fmt.Sprintf("no activation token: %s", activationtoken))
		return c.JSON(http.StatusInternalServerError, "error")
	}

	_, err := uuid.Parse(activationtoken)
	if err != nil {
		tracing.LogString(span, "activationtoken", fmt.Sprintf("illegal token: %s", activationtoken))
		c.Logger().Warnf("Illegal token activationtoken %s", activationtoken)
		return c.JSON(http.StatusInternalServerError, "error")
	}

	tracing.LogString(span, "request-activation-state", fmt.Sprintf("Get Activationstate for activationtoken %s", activationtoken))
	c.Logger().Debugf("Get Activationstate for activationtoken %s", activationtoken)

	//TODO - check if activation code is too old (aka invalid)

	activator := &AccountActivation{
		Token: activationtoken,
	}

	err = activator.activate(c)

	if err != nil {
		tracing.LogString(span, "activation-state-request-error", fmt.Sprintf("%v", err))
		return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf(serverConfig.Activation.AcitvationStateURL, "ongoing"))
	}
	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf(serverConfig.Activation.AcitvationStateURL, "success"))
}

// End of Account Activation stuff

// activateMerchant sets the merchant flag to true for logged in user.
func activateMerchant(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	return c.JSON(http.StatusOK, setMerchant(c, true))
}

// deactivateMerchant sets the merchant flag to false for logged in user.
func deactivateMerchant(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	return c.JSON(http.StatusOK, setMerchant(c, false))
}

// setMerchant set the given merchant flag for logged in user.
func setMerchant(c echo.Context, activate bool) MerchantModeResponse {
	span := tracing.Enter(c)
	defer span.Finish()

	response := MerchantModeResponse{Successful: false}
	sessionUser, err := findSessionUser(c)
	if err != nil {
		c.Logger().Info(err)
		tracing.LogError(span, err)
		return response
	}

	sessionUser.Merchant = activate
	err = sessionUser.Update(tracing.GetContext(c))
	if err != nil {
		c.Logger().Info(err)
		tracing.LogError(span, err)
		return response
	}

	response.Successful = true

	//TODO - Introduce i18n
	saveMessageRequest := &bubblesclient.SaveMessageRequest{
		Title:    "Merchant mode",
		UserID:   sessionUser.UUID,
		Category: bubblesclient.Categories[2],
		SenderID: bubblesclient.SenderIDs[1],
		Audience: bubblesclient.Audiences[0],
	}
	if activate {
		saveMessageRequest.Text = "successfully activated"
	} else {
		saveMessageRequest.Text = "successfully deactivated"
	}

	tracing.LogStruct(span, "message for receiver", saveMessageRequest)
	bubblesclient.SaveMessage(saveMessageRequest)
	return response
}

// getCallbackURL builds the entire callback-URL from callbackURL.
func getCallbackURL(c echo.Context, urlPart string) string {
	span := tracing.Enter(c)
	defer span.Finish()

	baseURL := serverConfig.Baseurl
	if urlPart == "" {
		tracing.LogString(span, "urlPart is empty, baseURL is used", baseURL)
		return baseURL
	}
	if urlPart[0] != '/' {
		callbackURL := baseURL + "/" + urlPart
		tracing.LogString(span, "callbackURL", callbackURL)
		return callbackURL
	}
	callbackURL := baseURL + urlPart
	tracing.LogString(span, "callbackURL (based on  urlPart with '/' only)", callbackURL)
	return callbackURL
}

// getCallbackURLFromSession returns URL-Part from cookie and deletes it after reading. Returning full callback-URL.
func getCallbackURLFromSession(c echo.Context) string {
	span := tracing.Enter(c)
	defer span.Finish()

	sess, _ := session.Get(sessionName, c)
	urlPart := ""
	if sess.Values["callbackURL"] != nil {
		urlPart = sess.Values["callbackURL"].(string)
		tracing.LogString(span, "from session-values 'urlPart'", urlPart)
	}
	callbackURL := getCallbackURL(c, urlPart)
	tracing.LogString(span, "callback-URL", callbackURL)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	return callbackURL
}

func isExpectedPostRequest(queryParams url.Values, httpMethod string) bool {
	if queryParams.Encode() == "" && httpMethod == "POST" {
		return true
	}
	return false
}
