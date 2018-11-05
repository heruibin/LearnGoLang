package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"Learn/DB/MongoDboDb/MongoByMgoWithPool"
)

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	JoinedAt  time.Time     `bson:"joned_at"`
	Interests []string      `bson:"interests"`
}

func main() {
	MongoByMgoWithPool.Init()
	session := MongoByMgoWithPool.CloneSession() //调用这个获得session
	defer session.Close()                        //一定要记得释放

	c := session.DB("music_user").C("people_pool")
	err := c.Insert(&User{Id_: bson.NewObjectId(),
		Name: "Jimmy Kuu",
		Age: 33,
		JoinedAt: time.Now(),
		Interests: []string{"Develop", "Movie"}})

	if err != nil {
		panic(err)
	}

	var users []User
	err = c.Find(nil).Limit(5).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

}
