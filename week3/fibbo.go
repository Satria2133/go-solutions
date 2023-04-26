package main

import "fmt"

func main() {
	var n int
	var fibbo int
	var nacci int
	fmt.Scan(&n)
	fibbo = 0
	nacci = 1
	for i := 0; i <= n; i++ {
		fmt.Print(fibbo, " ")
		fibbo = fibbo + nacci
		nacci = fibbo - nacci

	}
}
