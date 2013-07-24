package main

import "fmt"

func main() {
	temperatures := [6]float64{9.1, 9.6, 10.3, 10.5, 10.9, 11.3}
	var array [5]int

	fmt.Println("Temperatures every hour for the last six hours:")
	fmt.Printf("\t%v\n", temperatures)

	fmt.Printf("\n%d values (capacity: %d)\n", len(temperatures), cap(temperatures))
	fmt.Printf("\n%d values (capacity: %d)\n", len(array), cap(array))

	fmt.Printf("array: %v\n", array)
	array[1] = 1
	fmt.Printf("array: %v\n", array)
}
