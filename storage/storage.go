package storage

import (
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/app"
	"gopkg.in/mgo.v2"
)

const (
	collection = "user"
)

var (
	singleSession *mgo.Session
)

func init() {
	var err error
	singleSession, err = mgo.Dial(app.DatabaseStore)
	singleSession.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
	}
	Users = newMongoUserStore(singleSession)
}
