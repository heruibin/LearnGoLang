package main

import (
	"fmt"
	"strconv"
)

func main() {
	var aMap = initMap()
	for i := 0; i < 10 ; i++  {
		addData(aMap,"key" + strconv.Itoa(i),"value" + strconv.Itoa(i))
	}
	fmt.Printf("map:",aMap)
	delete(aMap,"key")
	fmt.Println()
	fmt.Printf("map:",aMap)
	fmt.Println()
	iteratorMap(aMap)
}

func initMap() map[string]string  {
	aMap := make(map[string]string)
	aMap["key"] = "value"
	return aMap
}

func addData(aMap map[string]string,key string,value string){
	aMap[key] = value;
}

func delData(aMap map[string]string,key string,value string){
	delete(aMap, "key")
}

func iteratorMap(aMap map[string]string){
	for key,value := range aMap{
		fmt.Printf("%s = %s",key,value)
		fmt.Println()
	}
}
