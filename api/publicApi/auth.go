package publicApi

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/app"
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/auth"
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/models"
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/storage"
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/templating"
	"github.com/labstack/echo"
	"github.com/markbates/goth/gothic"
)

type (
	signingProvider struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
)

var (
	tmplt *template.Template
	err   error
)

func init() {
	tmplt, err = templating.GetTemplate("oauthCallback.html")
}

// InitAuthWebServices Wire Auth Services ...
func InitAuthWebServices(app *echo.Echo) {
	gProvider := app.Group("/auth")
	gProvider.Get("", getProviders)
	gProvider.Get("/:provider", redirectHandler)
	gProvider.Get("/:provider/callback", callbackHandler)
}

func getProviders(c *echo.Context) error {
	providers := auth.GetProviders()
	results := []*signingProvider{}
	for _, provider := range providers {
		results = append(results, &signingProvider{
			provider.Name(),
			app.App.URI(redirectHandler, provider.Name()),
		})
	}
	return c.JSON(http.StatusOK, results)
}

func redirectHandler(c *echo.Context) error {
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

func callbackHandler(c *echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		log.Println(err)
		return err
	}
	provider, _ := gothic.GetProviderName(c.Request())
	u, err := storage.Users.Find(provider, user.UserID)
	if err != nil {
		u = &models.User{
			ID:        models.NewID(),
			Provider:  provider,
			Subject:   user.UserID,
			Name:      user.Name,
			Email:     user.Email,
			AvatarURL: user.AvatarURL,
			Role:      "readOnly",
			Active:    true,
			Created:   time.Now().UTC(),
		}
		storage.Users.Put(u)
	}

	log.Println("the provider is : ", provider)
	tmplt.ExecuteTemplate(c.Response(), "oauthCallback.html", user)

	return nil
}
