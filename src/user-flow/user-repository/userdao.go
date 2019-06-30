package user_repository

import (
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"user-flow/model"
)

var (
	session *mgo.Session
	c       *mgo.Collection
)

func Build(databaseURL string) {
	session, err := mgo.Dial(databaseURL)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c = session.DB("test").C("users")
}

func Close() {
	session.Close()
}

func Insert(user *model.User) *model.User {
	if len(user.ID) == 0 {
		user.ID = uuid.New().String()
	}
	err := c.Insert(*user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Inserted: ", user)

	return user
}

func Update(user *model.User) *model.User {
	err := c.Update(bson.M{"id": user.ID}, *user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Inserted: ", user)

	return user
}

func Delete(userID string) {
	user := &model.User{}
	err := c.Find(bson.M{"id": userID}).One(&user)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Remove(user)
	if err != nil {
		log.Fatal(err)
	}
}

func FindAll() (users []model.User, err error) {
	err = c.Find(nil).All(&users)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Find(userID string) (user model.User, err error) {
	err = c.Find(bson.M{"id": userID}).One(&user)
	if err != nil {
		log.Fatal(err)
	}
	return
}
