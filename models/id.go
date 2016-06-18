package models

import (
	"github.com/satori/go.uuid"
)

//NewID ...
func NewID() uuid.UUID {
	return uuid.NewV4()
}
