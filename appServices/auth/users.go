package auth

import (
	"time"

	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/models"
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/storage"
	"github.com/markbates/goth"
)

// CreateNewOAuthUser : create new logged-in user
func CreateNewOAuthUser(user goth.User) (*models.User, error) {
	u := &models.User{
		ID:        models.NewID(),
		Provider:  user.Provider,
		Subject:   user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		Role:      "readOnly",
		Active:    true,
		Created:   time.Now().UTC(),
	}
	return u, storage.Users.Put(u)
}
