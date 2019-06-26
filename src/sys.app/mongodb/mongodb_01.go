package mongodb

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type dbRecord struct {
	Number int    `bson:"number"`
	Text   string `bson:"text"`
}

// ConnectToDbAndTest : connect to mongo db, and test
func ConnectToDbAndTest() {
	addr := "127.0.0.1:27017"
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// query
	c := session.DB("zjdb").C("test")
	result := dbRecord{}
	if err = c.Find(bson.M{"number": 1}).One(&result); err == nil {
		fmt.Printf("number: %d, text: %s.\n", result.Number, result.Text)
	}

	// insert
	err = c.Insert(&dbRecord{3, "three"}, &dbRecord{4, "four"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert done.")
}
