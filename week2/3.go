package main

import "fmt"

func main() {
	var n int
	var waktu int
	var sum int
	var avg float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&waktu)
		sum += waktu
	}
	avg = float64(sum) / float64(n)
	fmt.Printf("%.2f", avg)
}
