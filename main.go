package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	j := Lexer(s)
	fmt.Println(j)
}

func Lexer(s string) []string {
	return strings.Split(s, " ")
}
