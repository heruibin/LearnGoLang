package main

import (
	"crypto/sha256"
	"fmt"
	"crypto/md5"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	fmt.Printf("%x", md5.Sum([]byte("12345")))

}
