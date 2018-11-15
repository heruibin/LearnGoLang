package main

import (
	tempconv "LearnGoLang/Base/Package/Standard"
	"fmt"
	"strconv"
	"log"
)

func main() {
	args := []string{"1", "2", "3"}
	for _, arg := range args {
		tmp, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		f := tempconv.Fahrenheit(tmp)
		c := tempconv.Celsius(tmp)
		fmt.Printf("%s = %s,%s = %s \n\a", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	var c tempconv.Celsius = 100
	fmt.Println(tempconv.CToF(c))
}
