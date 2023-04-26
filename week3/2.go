package main

import "fmt"

func main() {
	var kabisat int
	var isKabisat bool
	fmt.Scan(&kabisat)
	isKabisat = (kabisat%4 == 0 && kabisat%100 != 0) || kabisat%400 == 0
	fmt.Print(isKabisat)
}

/*
Program Kabisat
KD
	kabisat : integer
	isKabisat : boolean
Algoritma
	input kabisat
	isKabisat <--   kabisat mod 400 == 0 or   (kabisat mod 4 == 0 and kabisat mod 100 != 0)
	output isKabisat
endprogram
*/
