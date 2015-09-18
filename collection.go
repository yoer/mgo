package omgo

import "gopkg.in/mgo.v2"

type OmgoFunc func(*mgo.Collection) error

func RunMgFun(dbName string, collection string, fun OmgoFunc) error {
	session := getSession(dbName)
	defer session.Close()

	return fun(session.DB("").C(collection))
}
