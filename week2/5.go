package main

import "fmt"

func main() {
	var input int
	var count int = 0
	fmt.Scan(&input)
	for input != -1 {
		count += input
		fmt.Scan(&input)

	}
	fmt.Print(count >= 10)
}
