package main

import "fmt"

func main() {
	var input rune // char
	var isAlphanumeric bool
	input = '5'
	isAlphanumeric = (input >= 'a' && input <= 'z') || (input >= 'A' && input <= 'Z') || (input >= '0' && input <= '9')
	fmt.Print(isAlphanumeric)

}
