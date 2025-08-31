package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(agentaBold, "The MyRedis Database!", Normal)
	Hello()
}

func Hello() {
	fmt.Println("Hello")
	fmt.Println(ToUpper("Goodbye"))
	fmt.Println(strings.ToUpper("Goodbye"))
	fmt.Println(ToLower("Hello, World!"))
}

func ToLower(text string) string {
	const dist = 'a' - 'A'
	var sb strings.Builder
	for _, ch := range text {
		if ch >= 'A' && ch <= 'Z' {
			sb.WriteRune(ch + dist)
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func HelloHello() {
	fmt.Println("Hello, World!")
}

func ToUpper(text string) string {
	const dist = 'a' - 'A'
	var sb strings.Builder
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			sb.WriteRune(ch - dist)
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func AddNumbers() {
	fmt.Println("Hello World!")
}

func HelloWorld() {
	fmt.Println()
}
