package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const helloString = "Hello, OTUS!"

func main() {
	fmt.Print(stringutil.Reverse(helloString))
}
