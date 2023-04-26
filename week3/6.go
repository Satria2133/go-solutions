package main

import "fmt"

func main() {
	var num1, num2, num3 int
	var isIntersect bool
	fmt.Scan(&num1, &num2, &num3)
	isIntersect = !((num1+num2) > num3 && (num1+num3) > num2 && (num2+num3) > num1)
	fmt.Print(isIntersect)
}

/*
Program isIntersect
KD
	num1, num2, num3 : integer
	isIntersect : boolean
	input (num1, num2, num3)
	isIntersect = not((num1+num2) > num3 and (num1+num3) > num2 and (num2+num3) > num1)
	output (isIntersect)
*/
