package main

import (
	"fmt"
)

func main() {
	fmt.Println(MagentaBold, "The MyRedis Database!", Normal)
	//Hello()
}

func Hello() {
	fmt.Println("Hello")
	fmt.Println("Goodbye")
	fmt.Println(AddTwoNumber(20, 30))
	fmt.Println("")
}

func AddTwoNumber(a, b int) int {
	result := a + b
	return result
}

func HelloWorld() {
	fmt.Println()
}
