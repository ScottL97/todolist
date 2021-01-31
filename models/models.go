package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

var mgoSession *mgo.Session

type Item struct {
	Id       int      `form:"-"`
	Title    string   `form:"title"`
	Detail   string   `form:"detail"`
	Deadline string   `form:"deadline"`
	People   []string `form:"-"`
}

func init() {
	username := beego.AppConfig.String("mongodb::mongo_username")
	password := beego.AppConfig.String("mongodb::mongo_password")
	if password == "123456" {
		fmt.Println("WARNING!!!Please modify the default db password!!!")
	}
	host := beego.AppConfig.String("mongodb::mongo_ip") + ":" + beego.AppConfig.String("mongodb::mongo_port")
	diaInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Username: username,
		Password: password,
	}
	s, err := mgo.DialWithInfo(diaInfo)
	if err != nil {
		log.Fatalln("create mongodb session error:", err)
	}
	mgoSession = s
}
func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	session := mgoSession.Copy()
	c := session.DB(db).C(collection)
	return session, c
}
func Count(db, collection string) (int, error) {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(nil).Count()
}
func FindAll(db, collection string, query, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Find(query).All(result)
}
func FindAllWithSelect(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Find(query).Select(selector).All(result)
}
func FindOne(db, collection string, query, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Find(query).One(result)
}
func FindOneWithSelect(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Find(query).Select(selector).One(result)
}
func Insert(db, collection string, docs ...interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()

	return c.Insert(docs...)
}
func Remove(db, collection string, query interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Remove(query)
}
func Update(db, collection string, query, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Update(query, update)
}
