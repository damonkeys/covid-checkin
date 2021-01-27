package main

import (
	"github.com/damonkeys/ch3ck1n/monkeys/languagehelper"
	//"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Config contains all necessary configuration informations
type Config struct {
	Port string `env:"SERVER_PORT"`
}

func main() {
	e := echo.New()
	e.Pre(middleware.NonWWWRedirect())
	e.Use(middleware.Recover())
	e.GET("/", checkLangRedirect)
	defaultENBox := packr.NewBox("./static_en_default")
	enBoxHTTPHandler := http.FileServer(defaultENBox)
	e.GET("/en/*", echo.WrapHandler(http.StripPrefix("/en/", enBoxHTTPHandler)))

	deBox := packr.NewBox("./static_de")
	deBoxHTTPHandler := http.FileServer(deBox)
	e.GET("/de/*", echo.WrapHandler(http.StripPrefix("/de/", deBoxHTTPHandler)))
	e.Logger.Fatal(e.Start(":4444"))

}

func checkLangRedirect(ec echo.Context) error {
	languageHeader := ec.Request().Header.Get("Accept-Language")
	languageCode := languagehelper.Retrieve(languageHeader)
	if strings.EqualFold("de", languageCode) {
		return ec.Redirect(http.StatusFound, "/de/")
	}
	return ec.Redirect(http.StatusFound, "/en/")
}
