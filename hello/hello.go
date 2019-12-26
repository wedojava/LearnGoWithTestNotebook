package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloSuffix = "!"

func main() {
	fmt.Println(Hello("world"))
}

// Hello just a begin of Go programing
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name + englishHelloSuffix
}
