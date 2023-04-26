package main

import "fmt"

func main() {
	var jariJari float64
	var volume float64
	var pi float64 = 3.14
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&jariJari)
	volume = 4.0 * pi * jariJari * jariJari * jariJari / 3.0
	fmt.Print(volume)
}
