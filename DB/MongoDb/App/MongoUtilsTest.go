package main

import (
	"LearnGoLang/DB/MongoDb/MongoByMgoWithPool"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	MongoByMgoWithPool.Init()

	//var user interface{}

	//user = User{Id_: bson.NewObjectId(),
	//	Name: "Jimmy Kuu",
	//	Age: 33,
	//	JoinedAt: time.Now(),
	//	Interests: []string{"Develop", "Movie"}}
	//
	//MongoByMgoWithPool.Insert("music_user", "test_util", &user)

	//var datas map[string]interface{} = make(map[string]interface{})
	//datas["name"] = "heruibin"
	//datas["age"] = 28
	//
	//MongoByMgoWithPool.Insert("music_user", "test_util2", &datas)

	var dataArr []interface{} = make([]interface{}, 2)
	var dataArr0 map[string]interface{} = make(map[string]interface{})
	dataArr0["name"] = "rbhe"
	dataArr0["age"] = 3
	dataArr[0] = dataArr0
	var dataArr1 map[string]interface{} = make(map[string]interface{})
	dataArr1["name"] = "he"
	dataArr1["age"] = 4
	dataArr[1] = dataArr1

	var users []interface{} = make([]interface{}, 2)
	users[0] = User{Name: "rbhe", Age: 1}
	users[1] = User{Name: "ruibin", Age: 2}

	MongoByMgoWithPool.InsertBatch("music_user", "test_util3", dataArr)

	//var params map[string]interface{} = make(map[string]interface{})
	//params["name"] = "heruibin"
	//
	//var values map[string]interface{} = make(map[string]interface{})
	//values["age"] = 30
	//
	//var sets map[string]interface{} = make(map[string]interface{})
	//sets["$set"] = values
	//
	//MongoByMgoWithPool.UpdateAll("music_user", "test_util2", params, sets)

	//session := MongoByMgoWithPool.CloneSession() //调用这个获得session
	//defer session.Close()                        //一定要记得释放
	//
	//c := session.DB("music_user").C("people_pool")
	//err := c.Insert(&User{Id_: bson.NewObjectId(),
	//	Name: "Jimmy Kuu",
	//	Age: 33,
	//	JoinedAt: time.Now(),
	//	Interests: []string{"Develop", "Movie"}})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//var users []User
	//err = c.Find(nil).Limit(5).All(&users)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(users)

}
