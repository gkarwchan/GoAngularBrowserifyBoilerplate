package auth

import (
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/settings"
	"github.com/markbates/goth"
    "net/http"
	"strings"
	"errors"
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
	
	if settings.GplusKey != "" {
		useGPlus(settings.GplusKey, settings.GplusSecret)
	}

	if settings.FacebookKey != "" {
		useFacebook(settings.FacebookKey, settings.FacebookSecret)
	}
}

//GetProviders ...
func GetProviders() goth.Providers {
	return goth.GetProviders()
}



func useGPlus(providerKey, providerSecret string) {
	goth.UseProviders(
		gplus.New(providerKey, providerSecret, settings.EndPointURL+"auth/gplus/callback"),
	)
}

func useFacebook(providerKey, providerSecret string) {
	goth.UseProviders(
		facebook.New(providerKey, providerSecret, settings.EndPointURL+"auth/facebook/callback"),
	)
}
