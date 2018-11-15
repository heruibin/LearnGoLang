package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "abc,def,ghi,jkl";
	idx := strings.LastIndex(str, ",")
	str = str[0:idx]
	fmt.Println(str)

	idx = strings.Index(str, ",")
	str = str[0:idx]
	fmt.Println(str)

	str = strings.ToUpper(str)
	fmt.Println(str)
}
