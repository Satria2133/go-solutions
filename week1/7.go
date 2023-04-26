package main

import "fmt"

func main() {
	var a int
	var b int
	var temp int
	fmt.Scan(&a, &b)
	temp = a
	a = b
	b = temp
	fmt.Print(a, b)
}
