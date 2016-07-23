package storage

import (
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/models"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// Users ...
	Users UserStore
)

type (
	// UserStore ...
	UserStore interface {
		Get(id uuid.UUID) (*models.User, error)
		Find(provider, subject string) (*models.User, error)
		Put(user *models.User) error
		// Post(user *models.User) error
		// List() []*models.User
		// GetUserByEmail(email string) (*models.User, error)
		// GetUserByRecoveryCode(recoveryCode uuid.UUID) (*models.User, error)
	}
	mongoUserStore struct {
		session    *mgo.Session
		collection string
	}
)

func newMongoUserStore(session *mgo.Session) UserStore {
	m := &mongoUserStore{
		session,
		"user",
	}

	index := mgo.Index{
		Key:        []string{"provider", "subject"},
		Unique:     true,
		Background: true,
	}
	c := session.DB("").C("user")
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return m
}

func (m *mongoUserStore) Get(id uuid.UUID) (*models.User, error) {
	s := m.session.Clone()
	defer s.Close()

	c := s.DB("").C(collection)
	q := c.FindId(id)
	user := &models.User{}
	err := q.One(user)
	return user, err
}

func (m *mongoUserStore) Find(provider, subject string) (*models.User, error) {
	s := m.session.Clone()
	defer s.Close()

	c := s.DB("").C(m.collection)
	q := bson.M{"provider": provider, "subject": subject}

	user := &models.User{}
	err := c.Find(q).One(user)
	return user, err
}

func (m *mongoUserStore) Put(user *models.User) error {
	s := m.session.Clone()
	defer s.Close()

	c := s.DB("").C(m.collection)

	_, err := c.UpsertId(user.ID, user)
	return err
}
