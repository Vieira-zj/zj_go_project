package mongodb

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Oplog : op log results struct
type Oplog struct {
	Timestamp    bson.MongoTimestamp `bson:"ts"`
	HistoryID    int64               `bson:"h"`
	MongoVersion int                 `bson:"v"`
	Operation    string              `bson:"op"`
	Namespace    string              `bson:"ns"`
	Object       bson.M              `bson:"o"`
	QueryObject  bson.M              `bson:"o2"`
}

func lastTime(session *mgo.Session) bson.MongoTimestamp {
	var member Oplog
	session.DB("local").C("oplog.rs").Find(nil).Sort("-$natural").One(&member)
	return member.Timestamp
}

// PrintMongoOpLog : print mongodb op log
func PrintMongoOpLog() {
	session, err := mgo.Dial(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Eventual, true)
	log.Println("db session:", session)

	iter := session.DB("local").C("oplog.rs").Find(bson.M{"ts": bson.M{"$gt": lastTime(session)}}).LogReplay().Tail(time.Second * 5)
	var result interface{}
	for {
		for iter.Next(&result) {
			log.Printf("%+v\n", result)
		}
		if iter.Err() != nil {
			log.Fatal(iter.Err())
		}
		if iter.Timeout() {
			log.Println("query timeout")
		}
	}
}
