package main

import "fmt"

func main() {
	var belanja int
	var kartu bool
	var diskon bool
	var cashback bool
	var isKartu bool
	fmt.Scan(&belanja, &kartu)
	diskon = belanja >= 100000
	isKartu = diskon && kartu
	cashback = (belanja >= 200000) && isKartu
	fmt.Println("Kartu? ", isKartu)
	fmt.Println("diskon? ", diskon)
	fmt.Println("Cashback? ", cashback)

}
