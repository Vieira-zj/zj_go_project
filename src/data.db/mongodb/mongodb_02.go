package mongodb

import (
	"encoding/base64"
	"fmt"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const addrMongoBetaZ0 = "10.200.20.38:27017"

//
// tblmgr
//
type tblmgrRecord struct {
	Tbl      string `json:"tbl" bson:"tbl"`
	UID      uint32 `json:"uid" bson:"uid"`
	Itbl     uint32 `json:"itbl" bson:"itbl"`
	PhyTbl   string `json:"phy" bson:"phy"`
	Ctime    int64  `json:"ctime" bson:"ctime"`
	DropTime int64  `json:"drop" bson:"drop"`
	Region   string `json:"region" bson:"region"`
	Zone     string `json:"zone" bson:"zone"`
	Global   bool   `json:"global" bson:"global"`
	Line     bool   `json:"line" bson:"line"`

	Ouid  uint32 `json:"ouid" bson:"ouid,omitempty"`
	Oitbl uint32 `json:"oitbl" bson:"oitbl,omitempty"`
	Otbl  string `json:"otbl" bson:"otbl,omitempty"`
	Perm  uint32 `json:"perm" bson:"perm,omitempty"`
}

// QeuryFromTblmgr : query records for tblmgr
func QeuryFromTblmgr() {
	session, err := mgo.Dial(addrMongoBetaZ0)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	results := tblmgrRecord{}
	info := queryInfo{
		DataBase:   "qbox_rs",
		Collection: "tblmgr",
		Query:      bson.M{"tbl": "share_tbl_mkbucket_zXfTDx"},
		Results:    &results,
	}
	queryDBRecods(session, info)
	fmt.Printf("results: %+v\n", results)
}

//
// uc
//
type ucRecord struct {
	Key   string `json:"key" bson:"key"`
	Group string `json:"grp" bson:"grp"`
	Val   string `json:"val" bson:"val"`
}

// QeuryFromUc : query records for uc
func QeuryFromUc() {
	session, err := mgo.Dial(addrMongoBetaZ0)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	results := ucRecord{}
	info := queryInfo{
		DataBase:   "qbox_uc",
		Collection: "uc",
		Query:      bson.M{"tbl": "mtw8wd:KODO-3058-testUfmjXAWeag"},
		Results:    &results,
	}
	queryDBRecods(session, info)
	fmt.Printf("search results: %+v\n", results)
}

//
// pub
//
type pubRecord struct {
	Domain  string `json:"domain" bson:"domain"`
	Tbl     string `json:"tbl" bson:"tbl"`
	Owner   int64  `json:"owner" bson:"owner"`
	Refresh bool   `json:"refresh" bson:"refresh"`
	Global  bool   `json:"global" bson:"global"`
}

// QeuryFromPub : query records for pub domain
func QeuryFromPub() {
	session, err := mgo.Dial(addrMongoBetaZ0)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	results := pubRecord{}
	info := queryInfo{
		DataBase:   "qbox_pub",
		Collection: "pub",
		Query:      bson.M{"tbl": "KODO-3058-testUfmjXAWeag"},
		Results:    &results,
	}
	queryDBRecods(session, info)
	fmt.Printf("search results: %+v\n", results)
}

//
// bucket
//
type bucketRecord struct {
	Tbl    string `json:"tbl" bson:"tbl"`
	UID    uint32 `json:"uid" bson:"uid"`
	Itbl   uint32 `json:"itbl" bson:"itbl"`
	PhyTbl string `json:"phy" bson:"phy"`
	Ctime  int64  `json:"ctime" bson:"ctime"`
	// !=0时, 表示该条目被删除 (包括uc和domain)
	DropTime int64        `json:"drop" bson:"drop"`
	Region   string       `json:"region" bson:"region"`
	Zone     string       `json:"zone" bson:"zone"`
	Global   bool         `json:"global" bson:"global"`
	Line     bool         `json:"line" bson:"line"`
	Val      string       `json:"val" bson:"val,omitempty"`
	Domains  []domainInfo `json:"domain_info" bson:"domain_info,omitempty"`

	Ouid  uint32 `json:"ouid" bson:"ouid,omitempty"`
	Oitbl uint32 `json:"oitbl" bson:"oitbl,omitempty"`
	Otbl  string `json:"otbl" bson:"otbl,omitempty"`
	Perm  uint32 `json:"perm" bson:"perm,omitempty"`
}

type domainInfo struct {
	Domain  string `json:"domain" bson:"domain"`
	Refresh bool   `json:"refresh" bson:"refresh"`
	Global  bool   `json:"global" bson:"global"`
}

// QeuryFromBucket : query records for bucket
func QeuryFromBucket() {
	session, err := mgo.Dial(addrMongoBetaZ0)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	results := bucketRecord{}
	info := queryInfo{
		DataBase:   "qbox_bucket",
		Collection: "bucket",
		Query:      bson.M{"itbl": 512345379},
		Results:    &results,
	}
	queryDBRecods(session, info)
	fmt.Printf("search results: %+v\n", results)
}

type queryInfo struct {
	DataBase   string
	Collection string
	Query      interface{}
	Results    interface{}
}

func queryDBRecods(session *mgo.Session, info queryInfo) {
	c := session.DB(info.DataBase).C(info.Collection)
	err := c.Find(info.Query).One(info.Results)
	if err != nil {
		panic(err)
	}
}

//
// rs
//
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
