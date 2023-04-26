package main

import "fmt"

func main() {
	var jumlahKue int
	var anggotaKeluarga int
	var kueterSisa int
	fmt.Scan(&jumlahKue)
	fmt.Scan(&anggotaKeluarga)
	kueterSisa = jumlahKue % anggotaKeluarga
	fmt.Print(kueterSisa)
}
