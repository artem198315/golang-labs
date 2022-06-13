package main

import (
	"fmt"

	u "golang.org/x/example/stringutil"
)

var greetings = "Hello, OTUS!"

func main() {
	fmt.Println(u.Reverse(greetings))
}
