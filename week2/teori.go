package main

import "fmt"

func main() {
	/* Perulangan
	   for -->> perulangan berdasarkan jumlah, jadi di awal diketahui jumlah bakal berapa kali melakukan perulangan.
	   while -->> perulangan berdasarkan kondisi, akan terus berulang hingga kondisi tersebut tidak lagi memenuhi.
	   repeat until -->> perulangan berdasarkan kondisi, akan terus berulang hingga kondisi tersebut tidak lagi memenuhi.
	   iterasi -->> adalah hal yang membuat perulangan berjalan dan berhenti.
	   i++ -->> increment, menambahkan nilai i sebanyak 1.
	   i-- -->> decrement, mengurangi nilai i sebanyak 1.
	*/

	var nama string
	fmt.Scan(&nama)
	for nama != "doni" {
		fmt.Scan(&nama)
	}

	// var i int
	var num int = 5
	for i := 1; i <= num; i++ {
		fmt.Println("Perulangan ke", i)
	}
}
