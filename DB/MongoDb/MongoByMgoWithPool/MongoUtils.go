package MongoByMgoWithPool

import (
	"gopkg.in/mgo.v2"
	"time"
)

//global
var MgoSession *mgo.Session

type MongoCfg struct {
	Addr     string `yaml:"addr"`
	PoolSize int    `yaml:"poolSize"`
}

func Init(mongoCfg *MongoCfg) {
	globalMgoSession, err := mgo.DialWithTimeout(mongoCfg.Addr, 10*time.Second)
	if err != nil {
		panic(err)
	}
	MgoSession = globalMgoSession
	MgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	MgoSession.SetPoolLimit(mongoCfg.PoolSize)
}

func CloneSession() *mgo.Session {
	return MgoSession.Clone()
}

func Insert(dbName string, colName string, data *map[string]interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)
	err := c.Insert(data)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func InsertBatch(dbName string, colName string, datas []interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	bulk := c.Bulk()
	bulk.Unordered()
	bulk.Insert(datas...)
	_, err := bulk.Run()

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func FindOne(dbName string, colName string, params map[string]interface{}) map[string]interface{} {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	var result map[string]interface{}
	err := c.Find(params).Limit(1).All(&result)
	if err != nil {
		panic(err)
		return nil
	}
	return result
}

func FindByLimit(dbName string, colName string, params map[string]interface{}, limit int) []map[string]interface{} {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	var results []map[string]interface{}
	err := c.Find(params).Limit(limit).All(&results)
	if err != nil {
		panic(err)
		return nil
	}
	return results
}

func Update(dbName string, colName string, params map[string]interface{}, values map[string]interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	err := c.Update(params, values)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func UpdateAll(dbName string, colName string, params map[string]interface{}, values map[string]interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	_, err := c.UpdateAll(params, values)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func UpSert(dbName string, colName string, params map[string]interface{}, values map[string]interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	_, err := c.Upsert(params, values)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func DropCol(dbName string, colName string) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)
	err := c.DropCollection()

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func Remove(dbName string, colName string, params map[string]interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)
	err := c.Remove(params)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func RemoveAll(dbName string, colName string, params map[string]interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)
	_, err := c.RemoveAll(params)

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func RemoveBatch(dbName string, colName string, params []interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)

	bulk := c.Bulk()
	bulk.Remove(params...)
	_, err := bulk.Run()

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func RemovesBatchAll(dbName string, colName string, params []interface{}) bool {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)
	bulk := c.Bulk()
	bulk.Remove(params...)
	_, err := bulk.Run()

	if err != nil {
		panic(err)
		return false
	}
	return true
}

func Count(dbName string, colName string, params map[string]interface{}) int {
	session := CloneSession()
	defer session.Close()

	c := session.DB(dbName).C(colName)
	count, err := c.Find(params).Count()

	if err != nil {
		panic(err)
		return -1
	}
	return count
}
