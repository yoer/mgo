package omgo

import (
	"testing"
	"time"

	"github.com/yoer/omgo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mg_collection = "omgo_test_collection"
	mg_instance   = "mg_instance"
)

var (
	mg_user     = ""
	mg_pass     = ""
	mg_host     = "127.0.0.1:27017"
	mg_database = ""
)

func init() {
	omgo.AppendSession(mg_instance, &MgoDBCfg{
		User: mg_user,
		Pass: mg_pass,
		Host: mg_host,
		DB:   mg_database,
	})
}

//	golang structure in mongoDB
type DongoData struct {
	Id   bson.ObjectId       `bson:"_id"`
	Name string              `bson:`
	Date bson.MongoTimestamp `bson:"date"`
}

func TestInsert(t *testing.T) {
	if err := omgo.RunMgFun(mg_instance, mg_collection, func(c *mgo.Collection) error {
		return c.Insert(&DongoData{
			Id:   bson.NewObjectId(),
			Name: "yoer",
			Date: bson.MongoTimestamp(time.Now().UnixNano()),
		})
	}); nil != err {
		t.Error(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	if err := omgo.RunMgFun(mg_instance, mg_collection, func(c *mgo.Collection) error {
		return c.Update(bson.M{"name": "yoer"}, bson.M{"name": "yoer", "date": time.Now().UnixNano()})
	}); nil != err {
		t.Error(err.Error())
	}
}

func TestFind(t *testing.T) {
	find := &DongoData{}

	if err := omgo.RunMgFun(mg_instance, mg_collection, func(c *mgo.Collection) error {
		return c.Find(bson.M{"name": "yoer"}).One(find)
	}); nil != err {
		t.Error(err.Error())
	}

	t.Log(find)
}

func TestCount(t *testing.T) {
	count := 0

	if err := omgo.RunMgFun(mg_instance, mg_collection, func(c *mgo.Collection) error {
		rs, err := c.Find(bson.M{"name": "yoer"}).Count()

		count = rs
		return err
	}); nil != err {
		t.Error(err.Error())
	}

	t.Log(count)
}
