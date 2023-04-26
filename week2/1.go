package main

import "fmt"

func main() {
	var bilangan int
	var sebelumnya int
	var banyak int
	var hasil bool = true
	fmt.Scan(&bilangan)
	banyak = 0
	for hasil == true {
		sebelumnya = bilangan
		banyak++
		fmt.Scan(&bilangan)
		hasil = bilangan < sebelumnya && banyak < 10

	}
	fmt.Print(banyak)
}
