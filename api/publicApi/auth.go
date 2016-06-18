package publicApi

import (
	"net/http"
	"log"
	"html/template"
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/auth"
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/app"
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/templating"
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
	err error
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
	log.Printf("provider: %v", user.Provider)
	log.Printf("email: %v", user.Email)
	log.Printf("name: %v", user.Name)
	log.Printf("nickname: %v", user.NickName)
	log.Printf("desc: %v", user.Description)
	log.Printf("user id: %v", user.UserID)
	log.Printf("avatar url: %v", user.AvatarURL)
	log.Printf("location: %v", user.Location)
	log.Printf("token : %v", user.AccessToken)
	log.Printf("token secret: %v", user.AccessTokenSecret)
	log.Printf("refresh token: %v", user.RefreshToken)
	log.Printf("expire: %v", user.ExpiresAt)
	provider, _ := gothic.GetProviderName(c.Request())
	log.Println("the provider is : " , provider)
	tmplt.ExecuteTemplate(c.Response(), "oauthCallback.html", user)
	
	return nil
}
