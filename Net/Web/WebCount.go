package main

import (
	"sync"
	"net/http"
	"unsafe"
	"fmt"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", addCount)
	http.HandleFunc("/counter", qryCount)
	http.ListenAndServe("localhost:8081", nil)
}

func addCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Println(count, r.URL)
	mu.Unlock()
	w.Write([]byte("hello count"))
}

func qryCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func Int2Byte(data int) (ret []byte) {
	var len uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len)
	var tmp int = 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

func Byte2Int(data []byte) int {
	var ret int = 0
	var len int = len(data)
	var i uint = 0
	for i = 0; i < uint(len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}
