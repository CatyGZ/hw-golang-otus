package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	orinalText := "Hello, OTUS!"
	reverseText := stringutil.Reverse(orinalText)

	fmt.Println(reverseText)
}
