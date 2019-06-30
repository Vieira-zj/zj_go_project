package mongodb

import (
	"fmt"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MgoOperation mongodb general operations.
type MgoOperation struct {
	dbSession *mgo.Session
}

// Oplog mongodb op log record struct.
type Oplog struct {
	Timestamp    bson.MongoTimestamp `bson:"ts"`
	HistoryID    int64               `bson:"h"`
	MongoVersion int                 `bson:"v"`
	Operation    string              `bson:"op"`
	Namespace    string              `bson:"ns"`
	Object       bson.M              `bson:"o"`
	QueryObject  bson.M              `bson:"o2"`
}

// NewMgoOpertion return a MgoOperation instance.
func NewMgoOpertion(dbURI string) *MgoOperation {
	session, err := mgo.Dial(dbURI)
	if err != nil {
		panic(fmt.Sprintln("connect to mongodb failed:", err))
	}
	session.SetMode(mgo.Eventual, true)

	op := MgoOperation{
		dbSession: session,
	}
	return &op
}

// Close : close mongodb session.
func (op *MgoOperation) Close() {
	op.dbSession.Close()
}

// PrintMgoOpLogs print mongodb op logs.
func (op MgoOperation) PrintMgoOpLogs() {
	log.Println("db session:", op.dbSession)
	iter := op.dbSession.DB("local").C("oplog.rs").Find(
		bson.M{"ts": bson.M{"$gt": op.lastTime()}}).LogReplay().Tail(time.Second * 5)

	var result interface{}
	for {
		for iter.Next(&result) {
			log.Printf("%+v\n", result)
		}
		if iter.Err() != nil {
			log.Fatal(iter.Err())
		}
		if iter.Timeout() {
			log.Println("op log query timeout!")
		}
	}
}

func (op MgoOperation) lastTime() bson.MongoTimestamp {
	var member Oplog
	op.dbSession.DB("local").C("oplog.rs").Find(nil).Sort("-$natural").One(&member)
	return member.Timestamp
}
