package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/app"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
)

func init() {
	gothic.GetState = func(req *http.Request) string {
		// Get the state string from the query parameters.
		return req.URL.Query().Get("state")
	}
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		q := req.URL.Query()
		provider := q.Get("provider")
		if provider == "" {
			segments := strings.Split(req.URL.Path, "/")
			provider = segments[2]
		}
		if provider == "" {
			return provider, errors.New("you must select a provider")
		}
		return provider, nil
	}

	if app.GplusKey != "" {
		useGPlus(app.GplusKey, app.GplusSecret)
	}

	if app.FacebookKey != "" {
		useFacebook(app.FacebookKey, app.FacebookSecret)
	}
}

//GetProviders ...
func GetProviders() goth.Providers {
	return goth.GetProviders()
}

func useGPlus(providerKey, providerSecret string) {
	goth.UseProviders(
		gplus.New(providerKey, providerSecret, app.EndPointURL+"auth/gplus/callback"),
	)
}

func useFacebook(providerKey, providerSecret string) {
	goth.UseProviders(
		facebook.New(providerKey, providerSecret, app.EndPointURL+"auth/facebook/callback"),
	)
}
