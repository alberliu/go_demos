package mgo

import (
	"testing"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User3 struct {
	ID   bson.ObjectId `bson:"_id"`
	Name Name          `bson:"name"`
	Ege  int           `bson:"ege"`
	Book []int         `bson:"book"`
}

type Name struct {
	Firstname string `bson:"firstname"`
	Lastname  string `bson:"lastname"`
}

func initDB() (*mgo.Collection) {
	session, err := mgo.Dial("")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	dataBase := session.DB("test")

	return dataBase.C("user")

}

func TestMgo(t *testing.T) {
	user1 := User3{ID: bson.NewObjectId(), Name: Name{"1", "1"}, Ege: 1, Book: []int{1, 2}}
	user2 := User3{ID: bson.NewObjectId(), Name: Name{"2", "2"}, Ege: 2, Book: []int{3, 4}}
	user3 := User3{ID: bson.NewObjectId(), Name: Name{"3", "3"}, Ege: 3, Book: []int{5, 6}}

	err := initDB().Insert(user1, user2, user3)
	if err != nil {
		fmt.Println(err)
	}

}

func TestOne(t *testing.T) {
	user := new(User3)
	initDB().Find(bson.M{"_id": bson.M{"$lt": bson.ObjectIdHex("5b0d3f3252fc810c023c8ab9")}}).Sort("-_id").One(user)
	fmt.Println(user)
}

func TestNextOne(t *testing.T) {
	user := new(User3)
	initDB().Find(bson.M{"_id": bson.M{"$gt": bson.ObjectIdHex("5b0d3f3252fc810c023c8ab7")}}).Sort("_id").One(user)
	fmt.Println(user)
}

func TestAll(t *testing.T) {
	var users []User3
	err := initDB().Find(nil).Sort("-ege").Limit(3).All(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

func TestAllIn(t *testing.T) {
	var users []User3
	err := initDB().Find(bson.M{"name": []int{1, 2}}).All(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

func TestSelect(t *testing.T) {
	var user User3
	err := initDB().Find(bson.M{"name": "1"}).Select(bson.M{"ege": 1}).One(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", user)
}

func TestUpdate(t *testing.T) {
	err := initDB().Update(bson.M{"_id": bson.ObjectIdHex("5b1a339352fc810779c367ba")}, bson.M{"$set": bson.M{
		"name.firstname":"2"}})
	if err != nil {
		fmt.Println(err)
	}
}

// ObjectId("5b1a339352fc810779c367ba")

