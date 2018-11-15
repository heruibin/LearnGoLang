package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Hello World!"
	b := []byte(s)

	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("enc=[%s]\n", sEnc)

	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	} else {
		fmt.Printf("dec=[%s]\n", sDec)
	}
}
