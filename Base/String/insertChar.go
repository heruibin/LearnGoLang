package main

import (
	"fmt"
	"strings"
	"bytes"
	"strconv"
)

func main() {
	str := "abcdefghijklmnopqrst"
	str = comma(str)
	fmt.Println(str)

	rslt := strings.Contains(str, "lalala")
	fmt.Println(rslt)
	fmt.Println(strings.Count(str, "a"))
	fmt.Println(strings.Fields(str))
	fmt.Println(strings.HasPrefix(str, "ab"))
	fmt.Println(strings.Join(strings.Fields(str), "@"))

	str2 := "abc"
	fmt.Println(strings.Fields(str2))

	strs := []string{"a", "b", "c"}
	fmt.Println(strings.Join(strs, "@"))

	ints := []int{1, 2, 2, 3, 4, 5, 65}
	fmt.Println(intsToString(ints))

	str = "abcdefghijklmn";
	for _, ch := range str {
		fmt.Println(ch)
	}
	fmt.Println(comma2(str))
}

func comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return comma(str[: n-3]) + "," + str[n-3:]
}

func comma2(str string) string {
	var buf bytes.Buffer
	count := 0
	for _, ch := range str {
		if count != 0 && count%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(ch)
		count++
	}
	return buf.String()
}

func intsToString(ints []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, v := range ints {
		//if i > 0 {
		//	buf.WriteString(", ")
		//}
		//fmt.Fprintf(&buf, "%d", v)
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteByte(']')
	return buf.String()
}
