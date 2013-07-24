package main

import "fmt"

func main() {
	myString := "hello world."
	fmt.Printf("The last character in the string is '%s'\n",
	string(myString[len(myString) - 1]))
}
