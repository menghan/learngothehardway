package main

import "fmt"

const n = 1

func main() {
	var i = 1
	var flt = 1.0

	fmt.Printf("i = %v, flt = %v, n = %v\n", i, flt, n)
	fmt.Printf("i+%v = %d, flt+%v = %f\n", n, i + n, n, flt + n)
}
