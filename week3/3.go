package main

import "fmt"

func main() {
	var num1, num2, num3 int
	var isMidPoint bool
	fmt.Scan(&num1, &num2, &num3)
	isMidPoint = ((num1+num2)/2 == num3) || ((num1+num3)/2 == num2) || ((num2+num3)/2 == num1)
	fmt.Print(isMidPoint)
}

/*
Program isMidPoint
KD
	num1,num2,num3 : integer
Algo
	input num1,num2,num3
	isMidPoint <-- (num1+num2)/2 = num3 or (num1+num3)/2 = num2 or (num2+num3)/2 = num1
	output isMidPoint
endprogram
*/
