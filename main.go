package main 

import (
	"fmt"
	//"gopkg.in/couchbase/gocb.v1"
	"github.com/spf13/viper"
	"os"
	"github.com/bgonza2017/couchbasedemo/couchstore"
	"github.com/bgonza2017/couchbasedemo/types"
	"github.com/pkg/errors"
	gocb "gopkg.in/couchbase/gocb.v1"
	"strings"
)

const (
	ACCT_PREFIX = "acct"
	ORDER_PREFIX = "order"
	ORDER_PREFIX_TWO = "ordertwo"
	SERVICE_PREFIX = "service"
)

var (
	server couchstore.CouchServer
	err error
)

type User struct {
	Id string `json:"uid"`
	Email string `json:"email"`
	Interests []string `json:"interests"`
}

func NewCouchStore() (couchstore.CouchServer, error) {
	ds, err := couchstore.NewDBServer(couchstore.CouchConfig{
		ConnectionString:     viper.GetString("cbconnstr"),
		BucketName: viper.GetString("cbbucket"),
		BucketPassword:   os.Getenv("CBPASSWORD"),
	})
	if err != nil {
		return nil, errors.Wrap(err, "couchstore.NewCouchStore")
	}

	return ds, nil
}

func GetAccounts() {
	var bucket *gocb.Bucket 
	bucket = server.GetBucketUtil()
	fmt.Printf("%v\n\n", bucket)

	whr := fmt.Sprintf("where name like \"%s\"", "Asset")
	queryPieces := []string{"select * from admin", whr}
	sqry := strings.Join(queryPieces, " ")
	query := gocb.NewN1qlQuery(sqry)
	resultSet, _ := bucket.ExecuteN1qlQuery(query, nil)
	fmt.Printf("%v\n\n", resultSet)
	var row interface{}
	for resultSet.Next(&row) {
	    fmt.Printf("Results: %+v\n", row)
	}	

	accts := []types.Account{}
	server.GetDocuments(query, &accts)
	fmt.Printf("accts:%v\n\n", accts)
}

func AddAccount() {
	name := "00001"
	acct1 := types.Account{AccountID:name, Name:"Asset", Disabled:"false"}
	key := fmt.Sprintf("%s:%s", ACCT_PREFIX, name)
	server.UpsertDocument(key, &acct1, 0)	

	name = "20001"
	acct1 = types.Account{AccountID:name, Name:"Payables", Disabled:"false"}
	key = fmt.Sprintf("%s:%s", ACCT_PREFIX, name)
	server.UpsertDocument(key, &acct1, 0)	
}

func AddService() {
	item1 := types.Right{Name:"Read", Disabled:true}
	item2 := types.Right{Name:"Write", Disabled:false}
	item3 := types.Right{Name:"Add", Disabled:true}
	item4 := types.Right{Name:"Delete", Disabled:false}
	lineitems := []types.Right{item1, item2, item3, item4}
	//uuid1 := uuid.NewV4().String()
	name := "OrderEntrySystem"
	key := fmt.Sprintf("%s:%s", SERVICE_PREFIX, name)
	service := types.Service{
		ID:key,
		Name:name,
		Rights: lineitems,
	}

	server.UpsertDocument(key, &service, 0)	
}

func main() {
	os.Setenv("CBPASSWORD", "p8ssw0rd")
	viper.Set("cbconnstr", "localhost")
	viper.Set("cbbucket", "admin")

	server, err = NewCouchStore()

	if err != nil { 
	}

	if server == nil {
	}

	AddService()
	AddAccount()
	GetAccounts()

}