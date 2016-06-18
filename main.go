package main

import (
   "github.com/gkar68/GoAngularBrowserifyBoilerplate/api/publicApi"
   "github.com/gkar68/GoAngularBrowserifyBoilerplate/app"
   "github.com/gkar68/GoAngularBrowserifyBoilerplate/storage"
   "github.com/labstack/echo"
)


var (
	localApp *echo.Echo
)

func main() {
	localApp := app.CreateApp()
	publicApi.InitAuthWebServices(localApp)
	err := storage.InitDataStore()
	if err != nil {
		panic(err)
	}
	// messaging.WireSMTP()
	app.Run(localApp)
}
