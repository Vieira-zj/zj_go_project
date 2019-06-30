package mongodb

import (
	"encoding/base64"
	"fmt"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RSOperation db test operations for rs.
type RSOperation struct {
	tag       string
	dbSession *mgo.Session
}

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

// NewRsOperation return a RsOperation instance.
func NewRsOperation() *RSOperation {
	addr := "10.200.30.11:8001"
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(fmt.Sprintln("connect to mongodb failed:", err))
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	rs := RSOperation{
		tag:       "test",
		dbSession: session,
	}
	return &rs
}

// Close : close mongodb session.
func (rs *RSOperation) Close() {
	rs.dbSession.Close()
}

// InsertRsRecords do insert, query and delete operations for rs.
func (rs RSOperation) InsertRsRecords() {
	c := rs.dbSession.DB("beta_rs1").C("rs")

	// insert
	recordID := "zjtest:/zhengjin/data/rstest01"
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
	err := c.Insert(&record)
	if err != nil {
		panic(fmt.Sprintln("insert failed:", err))
	}

	// query
	result := rsRecord{}
	err = c.Find(bson.M{"_id": recordID}).One(&result)
	if err != nil {
		panic(fmt.Sprintln("query failed:", err))
	}
	fmt.Printf("query record: %+v\n", result)
	fmt.Println("file handle:", rs.base64Encode(result.FH))

	// delete
	err = c.Remove(bson.M{"_id": recordID})
	if err != nil {
		panic(fmt.Sprintln("delete failed:", err))
	}
}

// InsertToRsTbParallel : parallel insert 100W+ records in rs table.
func (rs RSOperation) InsertToRsTbParallel() {
	c := rs.dbSession.DB("beta_rs1").C("rs")

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
					fmt.Println(err)
				}
				if j%1000 == 0 {
					fmt.Printf("***** thread %d, insert %d records\n", idx, j)
				}
			}
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("parallel insert rs DONE.")
}

func (rs RSOperation) base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
