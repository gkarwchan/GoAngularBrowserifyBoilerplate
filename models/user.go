package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type (
	// User is the loggin user
	User struct {
		ID            uuid.UUID `json:"id"         bson:"_id"`
		Provider      string    `json:"provider"   bson:"provider"`
		Subject       string    `json:"subject"    bson:"subject"`
		Name          string    `json:"name"       bson:"name"`
		Email         string    `json:"email"      bson:"email"`
		Role          string    `json:"role"      bson:"role"`
		AvatarURL     string    `json:"avatar_url" bson:"avatar_url"`
		Active        bool      `json:"active"     bson:"active"`
		Created       time.Time `json:"created"    bson:"created"`
		Password      string    `json:"password"   bson:"-"`
		SavedPassword string    `json:"-"  bson:"password"`
	}
)
