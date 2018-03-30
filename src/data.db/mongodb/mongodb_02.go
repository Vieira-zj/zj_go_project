package mongodb

import (
	"encoding/base64"
	"fmt"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type fileRecord struct {
	ID       string `bson:"_id"`
	FDel     uint32 `bson:"fdel"`
	Hash     string `bson:"hash"`
	IP       []byte `bson:"ip"`
	FSize    int64  `bson:"fsize"`
	MimeType string `bson:"mimeType"`
	FH       []byte `bson:"fh"`
	PutTime  int64  `bson:"putTime"`
}

// InsertToRsDb : insert file records in kodo rs db
func InsertToRsDb() {
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
	record := fileRecord{
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
	result := fileRecord{}
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
			record := fileRecord{
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
