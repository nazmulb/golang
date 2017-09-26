package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		ID         bson.ObjectId `json:"id" bson:"_id"`
		Username   string        `json:"username" bson:"username"`
		Password   string        `json:"password" bson:"password"`
		Profession string        `json:"profession" bson:"profession"`
	}
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	checkErr(err)

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("myapp").C("users")

	// Insert
	err = c.Insert(&User{ID: bson.NewObjectId(), Username: "saiham", Password: "1234", Profession: "Senior Software Engineer"})
	checkErr(err)

	// Query One by username
	result := User{}
	err = c.Find(bson.M{"username": "saiham"}).One(&result)
	checkErr(err)

	fmt.Println("Profession:", result.Profession)

	// Query All by profession
	var results []User
	err = c.Find(bson.M{"profession": "Senior Software Engineer"}).Sort("-username").All(&results)
	checkErr(err)

	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"username": "saiham"}
	change := bson.M{"$set": bson.M{"profession": "DM"}}
	err = c.Update(colQuerier, change)
	checkErr(err)

	// Query All
	err = c.Find(bson.M{}).Sort("-username").All(&results)
	checkErr(err)

	fmt.Println("Results All: ", results)

	// Delete
	err = c.Remove(bson.M{"username": "saiham"})
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
