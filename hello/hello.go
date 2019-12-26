package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

// Hello just a begin of Go programing
func Hello(name string) string {
	return "Hello, " + name + "!"
}
