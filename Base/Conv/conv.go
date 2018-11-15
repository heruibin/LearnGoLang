package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	z := strconv.Itoa(x)
	j := strconv.FormatInt(int64(x), 8)
	fmt.Println(x, y, z, j)

	h, _ := strconv.Atoi("123")
	i, _ := strconv.ParseInt("3", 16, 64)

	fmt.Println(h, i)

}
