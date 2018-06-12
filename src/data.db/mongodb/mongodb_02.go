package mongodb

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//
// tblmgr, uc, pub => bucket
const (
	bucket = "test_bucket_transfer_data11"
	uid    = 1380469264
)

// QueryBucketInfo : query bucket info from db
func QueryBucketInfo() {
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
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	queryFromTblmgr(session)
	queryFromUc(session)
	queryFromPub(session)
	queryFromBucket(session)
	fmt.Println("query bucket info done.")
}

func queryFromTblmgr(session *mgo.Session) {
	info := queryInfo{
		DataBase:   "qbox_rs",
		Collection: "tblmgr",
		Query:      bson.M{"tbl": bucket},
	}
	results := queryDBOneRecord(session, info)
	fmt.Printf("query results: %+v\n", results)
}

func queryFromUc(session *mgo.Session) {
	key := strconv.FormatInt(uid, 36) + ":" + bucket
	info := queryInfo{
		DataBase:   "qbox_uc",
		Collection: "uc",
		Query:      bson.M{"key": key},
	}
	results := queryDBOneRecord(session, info)
	fmt.Printf("query results: %+v\n", results)
}

func queryFromPub(session *mgo.Session) {
	info := queryInfo{
		DataBase:   "qbox_pub",
		Collection: "pub",
		Query:      bson.M{"tbl": bucket},
	}
	results := queryDBAllRecords(session, info)
	fmt.Printf("query results: %+v\n", results)
}

func queryFromBucket(session *mgo.Session) {
	info := queryInfo{
		DataBase:   "qbox_bucket",
		Collection: "bucket",
		Query:      bson.M{"tbl": bucket},
	}
	results := queryDBOneRecord(session, info)
	fmt.Printf("query results: %+v\n", results)
}

type queryInfo struct {
	DataBase   string
	Collection string
	Query      interface{}
}

func queryDBOneRecord(session *mgo.Session, info queryInfo) interface{} {
	fmt.Printf("===> query one from %s.%s:\n", info.DataBase, info.Collection)
	c := session.DB(info.DataBase).C(info.Collection)
	num, err := c.Count()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("total:", num)

	var results interface{} // use interface{} instead of struct
	err = c.Find(info.Query).One(results)
	if err != nil {
		fmt.Println("error:", err)
	}
	return results
}

func queryDBAllRecords(session *mgo.Session, info queryInfo) []interface{} {
	fmt.Printf("===> query all from %s.%s:\n", info.DataBase, info.Collection)
	c := session.DB(info.DataBase).C(info.Collection)
	num, err := c.Count()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("total:", num)

	var results []interface{}
	err = c.Find(info.Query).All(&results)
	if err != nil {
		fmt.Println("error:", err)
	}
	return results
}

//
// rs
type rsRecord struct {
	ID       string `bson:"_id"`
	FDel     uint32 `bson:"fdel"`
	Hash     string `bson:"hash"`
	IP       []byte `bson:"ip"`
	FSize    int64  `bson:"fsize"`
	MimeType string `bson:"mimeType"`
	FH       []byte `bson:"fh"`
	PutTime  int64  `bson:"putTime"`
}

// InsertRsRecords : insert file records in kodo rs db
func InsertRsRecords() {
	addr := "10.200.30.11:8001"
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("beta_rs1").C("rs")

	recordID := "zjtest:/zhengjin/data/rstest01"
	// insert
	record := rsRecord{
		ID:       recordID,
		FDel:     1,
		FH:       []byte("rs-test-01"),
		FSize:    51523,
		Hash:     "FkBhdo9odL2Xjvu-YdwtDIw79fIX",
		IP:       []byte("172.21.6.102"),
		MimeType: "image/jpeg",
		PutTime:  15211684857206267,
	}
	err = c.Insert(&record)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert done.")

	// query
	result := rsRecord{}
	err = c.Find(bson.M{"_id": recordID}).One(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("results: %+v\n", result)
	fmt.Println("file handle:", base64Encode(result.FH))

	// delete
	err = c.Remove(bson.M{"_id": recordID})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("delete done.")
}

// InsertToRsDbParallel : insert 100W+ records in mongodb parallel
func InsertToRsDbParallel() {
	addr := "10.200.30.11:8001"
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("beta_rs1").C("rs")

	const (
		threads = 100 // max: 100
		count   = 10000
	)
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, idx int) {
			defer wg.Done()
			record := rsRecord{
				FDel:     0, // 1 - deleted
				FSize:    51523,
				Hash:     "FkBhdo9odL2Xjvu-YdwtDIw79fXX",
				IP:       []byte("172.21.6.102"),
				MimeType: "image/jpeg",
				PutTime:  15211684857206267,
			}
			for j := 0; j < count; j++ {
				// bucket: test_api_bucket_urts6VPcDs => 8h0ayi
				record.ID = fmt.Sprintf(fmt.Sprintf("8h0ayi:/zhengjin/dataE/test%de%d", idx, j))
				record.FH = []byte(fmt.Sprintf("rsAtestA%de%d", idx, j))
				err := c.Insert(&record)
				if err != nil {
					fmt.Println(err.Error())
				}
				if j%2000 == 0 {
					fmt.Printf("***** thread %d, insert %d records\n", idx, j)
				}
			}
		}(&wg, i)
	}
	wg.Wait()

	fmt.Println("insert done.")
}

func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
