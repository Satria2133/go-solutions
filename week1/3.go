package main

import "fmt"

func main() {
	var jumlahChakra int
	var butuhChakra int
	var klon int
	fmt.Scan(&jumlahChakra, &butuhChakra)
	klon = jumlahChakra / butuhChakra
	fmt.Print(klon)
}
