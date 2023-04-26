package main

import "fmt"

func main() {
	var jariJari int
	var tinggi int
	var volume float64
	var pi float64
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&jariJari)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)
	volume = 22 * pi * float64(jariJari) * float64(jariJari) * float64(tinggi) / 3.0 / 7.0
	fmt.Print(volume)
}


read <-> write
input <-> output
scan <-> print