package MongoByMgo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	//传入数据库的地址，可以传入多个，具体请看接口文档
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close() //用完记得关闭

	// Optional. Switch the session to a monotonic behavior.
	//读模式，与副本集有关，详情参考https://docs.mongodb.com/manual/reference/read-preference/ & https://docs.mongodb.com/manual/replication/
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("music_user").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result) //如果查询失败，返回“not found”
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
