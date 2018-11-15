package main

import (
	"bytes"
	"fmt"
	"AisporUtils/Collections/Set"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("abc")
	fmt.Println(buf.String())
	var buf2 bytes.Buffer
	buf2.WriteString("def")
	set := Set.New()
	set.Add(1)
	set.Add(1)
	fmt.Println(set.IsEmpty())
	for _, d := range set.List() {
		fmt.Println(d)
	}
}
