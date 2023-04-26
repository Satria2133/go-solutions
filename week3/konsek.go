package main

import "fmt"

func main() {

	var input int
	var digit1, digit2 int
	var selisih int
	var hasil bool = true

	fmt.Scan(&input)
	for input >= 10 && hasil {
		digit1 = input % 10
		digit2 = (input / 10) % 10
		selisih = digit1 - digit2
		input = input / 10
		hasil = selisih == 1 || selisih == -1
	}
	fmt.Print(hasil)
}

/*
scan || print
input || output
read || write

Program "JUDUL"
KD
	 num : integer
	 digit1, digit2 : integer
	 selisih : integer
	 hasil : boolean

Algoritma
	input num
	hasil <-- true
	while num >= 10 and hasil == true do
		digit1 <-- num mod 10
		digit2 <-- (num div 10) mod 10
		selisih <-- digit1 - digit2
		input <-- input div 10
		hasil <-- selisih == 1 or selisih == -1
	endwhile
	output hasil
*/
