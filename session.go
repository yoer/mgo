package omgo

import "gopkg.in/mgo.v2"

var (
	MgoSessions = map[string]*mgo.Session{}
)

func AppendSession(mg_instance string, config *MgoDBCfg) {
	MgoSessions[mg_instance] = createSession(getMgUrl(config))
}

func createSession(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err) //直接终止程序运行
	}

	return session
}

func getSession(mg_instance string) *mgo.Session {
	session, exist := MgoSessions[mg_instance]
	if !exist {
		panic("mongo db session is not exist, please call AppendSession!")
	}

	return session.Clone()
}
