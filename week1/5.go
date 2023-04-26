package main

import "fmt"

func main() {
	var panjang float64
	var lebar float64
	var luas float64
	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	fmt.Print(luas)
}
