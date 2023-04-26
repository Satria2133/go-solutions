package main

import "fmt"

func main() {
	var input int
	var hasil bool
	fmt.Scan(&input)
	// input bakal dicheck apakah dia negatif ganjil / bukan
	hasil = input < 0 && input%2 == 1 // input % 2 == 1
	fmt.Println(hasil)
}
