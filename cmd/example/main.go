package main

import (
	go21 "GO21"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Build version: " + version)
	input := strings.Join(os.Args[1:], " ")
	res, _ := go21.PostfixToInfix(input)
	fmt.Println(res)
}
