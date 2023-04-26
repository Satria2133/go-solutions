package main

import "fmt"

func main() {
	var input float64
	var sum float64 = 1
	fmt.Scan(&input)
	for input != 0 {
		sum *= (1 / input)
		fmt.Scan(&input)
	}
	fmt.Printf("%.4f", sum)
}
