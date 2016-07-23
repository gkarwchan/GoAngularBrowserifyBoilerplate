package main

import (
	"log"

	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/api/publicApi"
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/app"
	"github.com/labstack/echo"
)

var (
	localApp *echo.Echo
)

func main() {
	log.Println("strt .........")
	localApp := app.CreateApp()
	publicApi.InitAuthWebServices(localApp)
	// err := storage.InitDataStore()
	// if err != nil {
	// 	panic(err)
	// }
	// messaging.WireSMTP()
	app.Run(localApp)
}
