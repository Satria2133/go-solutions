package main

import "fmt"

func main() {
	var n int
	var input int
	var sum float64 = 0
	var HarmoniAvg float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		sum += (1 / float64(input))
	}
	HarmoniAvg = float64(n) / sum
	fmt.Print("%.2f", HarmoniAvg)
}
