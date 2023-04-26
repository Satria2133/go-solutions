package main

import "fmt"

func main() {
	var n int
	var m int
	var sum int = 0
	var sum0 int = 1
	fmt.Scan(&n, &m)
	for n <= m {
		sum += n
		n++
	}
	fmt.Print(sum)
}
