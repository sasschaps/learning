package main

import "fmt"

// Print out: double of 50
func main() {
	number := 50
	answer := double(number)

	fmt.Println(answer)
}

// double returns the number num multiplied by 2.
func double(num int) int {
	answer := num * 2

	return answer
}
