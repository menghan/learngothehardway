package main

import "fmt"

func main() {
	myString := "hello"
	fmt.Println("my string:", myString)

	myString = myString + " world"
	fmt.Println("my string:", myString)

	myString += "."
	fmt.Println("my string:", myString)

	fmt.Printf("my string: %d characters long\n", len(myString))

	fmt.Printf("the first character of my string is: %s\n", string(myString[0]))

	// fmt.Println("punctuation: ", punc)
}
