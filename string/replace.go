package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "123 123 123   123"
	str = strings.Replace(str, " ", "", -1)
	fmt.Println(str)
	unicode.ToLower(rune(str[1]))
	str = "Hello World 123 "
	str = strings.ToLower(str)
	fmt.Println(str)
	subStr := str[1:5]
	fmt.Println(subStr)
}
