package main

import (
	"fmt"
)

func main() {
	fmt.Println(MagentaBold, "The MyRedis Database!", Normal)
}

func Hello() {
	fmt.Println("Hello")
	fmt.Println("Goodbye")
	AddTwoNumber(20, 30)
}

func AddTwoNumber(a, b int) int {
	result := a + b
	fmt.Println(result)
	return result
}
