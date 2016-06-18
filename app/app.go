package app

import (
	"fmt"
	"log"

	"github.com/gkar68/GoAngularBrowserifyBoilerplate/settings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// App is going to be referenced by the web app and services
	App = initApp()
)

func initApp() *echo.Echo {
	// create router and stats handler
	e := echo.New()

	// why not, it's fast and "the future" ...
	e.HTTP2(true)

	// add middleware
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.Gzip())
	return e
}

// CreateApp ...
func CreateApp() *echo.Echo {
	var staticAssets = settings.StaticAssets
	App.Favicon(staticAssets + "/favicon.ico")
	App.Index(staticAssets + "/index.html")
	App.Static("/fonts", staticAssets+"/fonts")
	App.Static("/images", staticAssets+"/images")
	App.Static("/scripts", staticAssets+"/scripts")
	App.Static("/styles", staticAssets+"/styles")
	App.Static("/app", staticAssets+"/app")
	App.Static("/templates.js", staticAssets+"/templates.js")
	return App
}

// Run ...
func Run(application *echo.Echo) {
	var address = fmt.Sprintf("%s:%d", settings.WebHost, settings.WebPort)
	log.Printf("starting webserver on %s", address)
	if settings.UseSSL {
		application.RunTLS(address, settings.CertFile, settings.KeyFile)
	} else {
		application.Run(address)
	}
}
