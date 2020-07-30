package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/damonkeys/ch3ck1n/auth/models"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"

	"github.com/eefret/gravatar/default_img"

	"github.com/eefret/gravatar"
	"github.com/google/uuid"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// sessionOptions for session token
var sessionOptions = sessions.Options{
	Path:     "/",
	MaxAge:   86400 * 7,
	HttpOnly: true,
}

// createSession cookies creates a new cookie and adds or find the user from known users
func createNewSessionCookie(c echo.Context, user goth.User) error {
	span := tracing.Enter(c)
	defer span.Finish()

	// Save in session
	sess, _ := session.Get("_monkeycash_session", c)
	sess.Options = &sessionOptions

	c.Logger().Debugf("email from goth-user: %s", user.Email)
	tracing.LogString(span, "email from goth-user", user.Email)

	// new user?
	sessionUser, err := models.FindUserByEmail(tracing.GetContext(c), user.Email)
	if err != nil {
		c.Logger().Debug("ERROR at findUserByEmail!!!")
		c.Logger().Error(err)
		c.Logger().Info(err)
		tracing.LogError(span, err)
		sessionUser, err = createNewSessionUser(c, sessionUser, user)
		if err != nil {
			return err
		}
	}
	c.Logger().Debugf("monkeycash-user with email [%s] found", sessionUser.Email)
	tracing.LogStruct(span, "monkeycash-user", sessionUser)

	// existing user
	// does it have a new provider?
	sessionProvider, err := models.FindUserProviderByName(tracing.GetContext(c), *sessionUser, user.Provider)
	if err != nil {
		// no provider found, create a new one
		sessionProvider = createNewSessionProvider(c, sessionUser, user)
	}
	tracing.LogStruct(span, "provider data", sessionProvider)
	updateAvatar(c, sessionUser, sessionProvider)

	// todo tracing
	if sessionUser.Active {
		// Save user-id to session
		sessions.NewCookie("userid", sessionUser.UUID, sess.Options)
		sess.Values["userid"] = sessionUser.UUID
		sess.Save(c.Request(), c.Response())

		// build Callback-URL
		return c.Redirect(http.StatusTemporaryRedirect, getCallbackURLFromSession(c))
	}

	//non active User
	if sessionUser.ActivationToken != "" {
		//create token and send email
		userInfo := UserInfo{
			Email: sessionUser.Email,
		}
		err := userInfo.createAndStoreActivationTokenForUser(c)

		if err != nil {
			// Needs good error Page with Support or even better: telemetry alerts
			tracing.LogStruct(span, "activation token cannot be created", userInfo)
			return c.JSON(http.StatusInternalServerError, "error")
		}

		preparer := &CTAMailContext{
			templatename: "cta-tpl",
			recipient:    "support@chckr.de",
			sender:       "support@chckr.de",
			subject:      fmt.Sprintf("Zugangs-Aktivierung für %s", userInfo.Email),
			cta:          "Zugangs-Aktivierung notwendig",
			body:         "Sie haben es fast geschafft: Klicken Sie auf den Aktivierungslink um Die Einrichtung Ihrs Zugangs abzuschließen. Vielen Dank",
			ctalink:      fmt.Sprintf(serverConfig.Activation.URL, userInfo.ActivationToken),
			linktext:     "Jetzt klicken und aktivieren",
		}

		tracing.LogStruct(span, "activation-mail-data", preparer)
		sendMail(preparer)

		return c.Redirect(http.StatusTemporaryRedirect, "/activation/missing")
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/activation/start")
}

func createNewSessionUser(c echo.Context, sessionUser *models.User, user goth.User) (*models.User, error) {
	span := tracing.Enter(c)
	defer span.Finish()

	// User already exist but deleted?
	sessionUser, err := models.FindDeletedUserByEmail(tracing.GetContext(c), user.Email)
	if err == nil {
		// Soft-deleted user exists, undelete it
		err = sessionUser.Undelete(tracing.GetContext(c))
		if err != nil {
			return nil, err
		}
		return sessionUser, nil
	}

	// new user, new provider!
	sessionProvider := createNewProviderData(c, user)

	sessionUser.Email = user.Email
	sessionUser.UUID = uuid.New().String()
	sessionUser.Name = user.Name
	sessionUser.AvatarURL = user.AvatarURL
	sessionUser.ActivationToken = uuid.New().String()
	sessionUser.Active = false
	sessionUser.ActivationTokenCreation = time.Now()
	gravatarURL, err := fetchGravatarURLIfAvailable(c, user.Email)
	if err == nil {
		sessionUser.AvatarURL = gravatarURL
	}
	// err = models.CreateUser(tracing.GetContext(c), sessionUser)
	err = sessionUser.Create(tracing.GetContext(c))
	if err != nil {
		return nil, err
	}

	sessionUser.AppendProviderToUser(tracing.GetContext(c), sessionProvider)
	tracing.LogStruct(span, "created provider data", sessionProvider)
	return sessionUser, nil
}

func createNewSessionProvider(c echo.Context, sessionUser *models.User, user goth.User) *models.Provider {
	span := tracing.Enter(c)
	defer span.Finish()

	sessionUser.AvatarURL = user.AvatarURL
	sessionUser.Update(tracing.GetContext(c))
	sessionProvider := createNewProviderData(c, user)
	sessionUser.AppendProviderToUser(tracing.GetContext(c), sessionProvider)
	return sessionProvider
}

func updateAvatar(c echo.Context, sessionUser *models.User, sessionProvider *models.Provider) {
	span := tracing.Enter(c)
	defer span.Finish()

	// Already set AvatarURL?
	if sessionUser.AvatarURL == "" {
		sessionUser.AvatarURL = sessionProvider.AvatarURL
		sessionUser.Update(tracing.GetContext(c))
	}
}

// removeSessionCookie deletes the existing session-cookie
func removeSessionCookie(c echo.Context) error {
	span := tracing.Enter(c)
	defer span.Finish()

	sess, _ := session.Get("_monkeycash_session", c)
	sess.Options.MaxAge = -1
	err := sess.Save(c.Request(), c.Response())
	return err
}

// findSession validates session and returns an error or the user
func findSessionUser(c echo.Context) (*models.User, error) {
	span := tracing.Enter(c)
	defer span.Finish()

	sess, err := session.Get("_monkeycash_session", c)
	if err != nil || len(sess.Values) == 0 {
		// session not valid
		sessions.NewCookie("userid", "", sess.Options)
		sess.Values["userid"] = ""
		return new(models.User), errors.New("no valid session found")
	}
	// Find user by user-id
	sessionUser, err := models.FindUserByUUID(tracing.GetContext(c), sess.Values["userid"].(string))
	if err != nil {
		resultError := fmt.Errorf("User with userid %s not found - %s", sess.Values["userid"], err)
		sessions.NewCookie("userid", "", sess.Options)
		sess.Values["userid"] = ""
		return new(models.User), resultError
	}
	return sessionUser, nil
}

// createNewProviderData creates a data object only in memory for storing in database via the user
func createNewProviderData(c echo.Context, user goth.User) *models.Provider {
	span := tracing.Enter(c)
	defer span.Finish()

	providerData := new(models.Provider)
	providerData.ProviderName = user.Provider
	providerData.AccessToken = user.AccessToken
	providerData.RefreshToken = user.RefreshToken
	providerData.Name = user.Name
	providerData.Lastname = user.LastName
	providerData.Firstname = user.FirstName
	providerData.UserID = user.UserID
	providerData.AvatarURL = user.AvatarURL
	providerData.Nickname = user.NickName
	providerData.Location = user.Location
	providerData.ExpiresAt = user.ExpiresAt
	return providerData
}

func fetchGravatarURLIfAvailable(c echo.Context, email string) (string, error) {
	span := tracing.Enter(c)
	defer span.Finish()

	g, err := gravatar.New()
	if err != nil {
		log.Printf("Couldn't check gravatar for email %s. Error: %s", email, err)
		tracing.LogError(span, err)
		return "", err
	}
	g.SetSize(uint(1280))
	g.SetDefaultImage(default_img.DefaultImage.HTTP_404)
	gravatarURL := g.URLParse(email)
	gravatarImageData, err := g.Download(email)
	if err != nil {
		log.Printf("Error during gravatar image download. Maybe %s has no gravatar? Error: %s", email, err)
		tracing.LogError(span, err)
		return "", err
	}
	if len(gravatarImageData) > 0 {
		log.Printf("Found gravatar image data for email %s", email)
		tracing.LogError(span, err)
		return gravatarURL, nil
	}
	log.Printf("No gravatar found for email %s, returning error", email)
	return "", errors.New("no gravatar data found")
}
