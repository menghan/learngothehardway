package main

import "fmt"

func main() {
	temperatures := []float64{9.1, 9.6, 10.3, 10.5, 10.9, 11.3}

	fmt.Println("Temperatures every hour for the last six hours:")
	fmt.Printf("\t%v\n", temperatures)
	fmt.Printf("\nTemperatures: %d values (capacity: %d)\n", len(temperatures), cap(temperatures))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("\nslice: %d values (capacity: %d)\n", len(slice), cap(slice))

	slice = append(slice, 6)
	fmt.Printf("\nslice: %d values (capacity: %d)\n", len(slice), cap(slice))
}
