package storage

import (
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/models"
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/settings"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "user"
)

var (
	singleSession *mgo.Session
)

// InitDataStore ...
func InitDataStore() error {
	var err error
	singleSession, err = mgo.Dial(settings.DatabaseStore)
	singleSession.SetMode(mgo.Monotonic, true)
	if err != nil {
		return err
	}
	return nil
}

// StoreUser ...
func StoreUser(user *models.User) error {
	session := singleSession.Clone()
	defer session.Close()
	c := session.DB("").C(collection)
	err := c.Insert(user)
	if err == nil {
		q := bson.M{"provider": "email", "email": user.Email}
		user = &models.User{}
		err = c.Find(q).One(user)
		return err
	}
	return err
}

// GetUserByEmail ...
func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	session := singleSession.Clone()
	defer session.Close()

	c := session.DB("").C(collection)
	q := bson.M{"provider": "email", "email": email}
	err := c.Find(q).One(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//Get ...
func Get(id uuid.UUID) (*models.User, error) {
	session := singleSession.Clone()
	defer session.Close()
	c := session.DB("").C(collection)
	q := c.FindId(id)
	user := &models.User{}
	err := q.One(user)
	return user, err
}

//Put ...
func Put(user *models.User) error {
	s := singleSession.Clone()
	defer s.Close()

	c := s.DB("").C(collection)

	_, err := c.UpsertId(user.ID, user)
	return err
}
