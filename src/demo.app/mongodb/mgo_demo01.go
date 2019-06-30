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

// ConnectToDbAndTest : connect to mongo db, and test op
func ConnectToDbAndTest() {
	// create connection
	addr := "127.0.0.1:27017"
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(fmt.Sprintln("connect to mongodb failed:", err))
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// query
	const (
		databaseName  = "zjdb"
		collecionName = "test"
	)
	c := session.DB(databaseName).C(collecionName)
	result := dbRecord{}
	err = c.Find(bson.M{"number": 1}).One(&result)
	if err != nil {
		panic(fmt.Sprintln("query failed:", err))
	}
	fmt.Printf("query results: %d=%s\n", result.Number, result.Text)

	// insert
	err = c.Insert(&dbRecord{3, "three"}, &dbRecord{4, "four"})
	if err != nil {
		panic(fmt.Sprintln("insert failed:", err))
	}
	fmt.Println("insert success")
}
