package main

import "fmt"

func main() {
	var sisi1, sisi2, sisi3 int
	var isSegitiga bool
	fmt.Scan(&sisi1, &sisi2, &sisi3)
	isSegitiga = (sisi1+sisi2 > sisi3 && sisi1+sisi3 > sisi2 && sisi2+sisi3 > sisi1)
	fmt.Print(isSegitiga)
}

/*
Program isSegitiga
KD
	sisi1,sisi2,sisi3 : integer
	isSegitiga : boolean
Algo
	input sisi1,sisi2,sisi3
	isSegitiga <-- not(sisi1+sisi2 <= sisi3 or sisi1+sisi3 <= sisi2 or sisi2+sisi3 <= sisi1)
	output isSegitiga
endprogram

*/
