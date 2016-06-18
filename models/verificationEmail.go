package models

import (
	"github.com/satori/go.uuid"
)

type (
	// VerificationEmail is the loggin user
	VerificationEmail struct {
		UserID          uuid.UUID
		Email           string
		AppURL          string
		VerificationURL string
	}
)
