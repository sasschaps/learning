package main

import (
	"fmt"
	"strings"
)

// Print out: Hello SARAH
func main() {
	s := "Sarah" // This thing (type) is a string, with name of s, and value of Sarah
	c := capitalise(s)

	fmt.Println("Hello " + c)
}

func capitalise(word string) string {
	cword := strings.ToUpper(word)

	return cword
}
