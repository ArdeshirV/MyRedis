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
}

func AddTwoNumber(a, b int) int {
	result := a + b
	return result
}

func AddNumbers() {
	fmt.Println("Hello World!")
}

func HelloWorld() {
	fmt.Println()
}
