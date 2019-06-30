package mongodb

import (
	"fmt"
	"os"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BucketInfo get bucket info => tblmgr,uc,pub => bucket
type BucketInfo struct {
	bucket    string
	uid       int64
	dbSession *mgo.Session
}

type queryInfo struct {
	DataBase   string
	Collection string
	Query      interface{}
}

// NewBucketInfo return a BucketInfo instance.
func NewBucketInfo(bucket string, uid int64) *BucketInfo {
	return &BucketInfo{bucket: bucket, uid: uid}
}

// QueryBucketInfo get bucket info from db.
func (b *BucketInfo) QueryBucketInfo() {
	const (
		addrMongoBetaZ0 = "10.200.20.38:27017"
		addrMongoDevZ0  = "10.200.20.23:27017"
	)

	dbURI := addrMongoBetaZ0
	if os.Getenv("TEST_ENV") == "dev" {
		dbURI = addrMongoDevZ0
	}
	session, err := mgo.Dial(dbURI)
	if err != nil {
		panic(fmt.Sprintln("connect to mongodb failed:", err))
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	b.dbSession = session

	b.queryFromTblmgr()
	b.queryFromUc()
	b.queryFromPub()
	b.queryFromBucket()
	fmt.Println("query bucket info DONE.")
}

func (b BucketInfo) queryFromTblmgr() {
	info := queryInfo{
		DataBase:   "qbox_rs",
		Collection: "tblmgr",
		Query:      bson.M{"tbl": b.bucket},
	}
	results := b.queryOneRecordFromTb(info)
	fmt.Printf("tblmgr info: %+v\n", results)
}

func (b BucketInfo) queryFromUc() {
	key := strconv.FormatInt(b.uid, 36) + ":" + b.bucket
	info := queryInfo{
		DataBase:   "qbox_uc",
		Collection: "uc",
		Query:      bson.M{"key": key},
	}
	results := b.queryOneRecordFromTb(info)
	fmt.Printf("uc info: %+v\n", results)
}

func (b BucketInfo) queryFromPub() {
	info := queryInfo{
		DataBase:   "qbox_pub",
		Collection: "pub",
		Query:      bson.M{"tbl": b.bucket},
	}
	results := b.queryAllRecordsFromTb(info)
	fmt.Printf("pub info: %+v\n", results)
}

func (b BucketInfo) queryFromBucket() {
	info := queryInfo{
		DataBase:   "qbox_bucket",
		Collection: "bucket",
		Query:      bson.M{"tbl": b.bucket},
	}
	results := b.queryOneRecordFromTb(info)
	fmt.Printf("bucket info: %+v\n", results)
}

func (b BucketInfo) queryOneRecordFromTb(info queryInfo) interface{} {
	fmt.Printf("===> query one record from %s.%s:\n", info.DataBase, info.Collection)
	c := b.dbSession.DB(info.DataBase).C(info.Collection)
	count, err := c.Count()
	if err != nil {
		panic(fmt.Sprintln("get count failed:", err))
	}
	fmt.Println("collection records count:", count)

	var results interface{} // use interface{} instead of struct
	if count > 0 {
		err = c.Find(info.Query).One(&results)
		if err != nil {
			panic(fmt.Sprint("query failed:", err))
		}
	}
	return results
}

func (b BucketInfo) queryAllRecordsFromTb(info queryInfo) []interface{} {
	fmt.Printf("===> query all records from %s.%s:\n", info.DataBase, info.Collection)
	c := b.dbSession.DB(info.DataBase).C(info.Collection)
	count, err := c.Count()
	if err != nil {
		panic(fmt.Sprintln("get count failed:", err))
	}
	fmt.Println("collection records count:", count)

	var results []interface{} // use slice
	if count > 0 {
		err = c.Find(info.Query).All(&results)
		if err != nil {
			panic(fmt.Sprint("query failed:", err))
		}
	}
	return results
}
