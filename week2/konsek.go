package main

import "fmt"

func main() {
	var digit1, digit2 int
	var num int
	var hasil bool = true
	var selisih int
	fmt.Scan(&num)
	for num >= 10 && hasil {
		digit1 = num % 10
		digit2 = (num / 10) % 10
		selisih = digit1 - digit2
		num = num / 10
		hasil = selisih == 1 || selisih == -1
	}
	fmt.Print(hasil)
}
