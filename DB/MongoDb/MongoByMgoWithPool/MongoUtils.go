package MongoByMgoWithPool

import (
	"gopkg.in/mgo.v2"
	"time"
)
//global
var GlobalMgoSession *mgo.Session

func Init() {
	globalMgoSession, err := mgo.DialWithTimeout("localhost:27017", 10 * time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession=globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(300)
}

func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

func Insert(dbName string,colName string){

}