package main

import "fmt"

func main() {
	var n, i int
	var item1, item2, item3, item4, item5, valid bool
	fmt.Scan(&n)
	i = 0
	fmt.Scan(&item1, &item2, &item3, &item4, &item5)
	valid = item1 == true && item2 == true && item3 == true && item4 == true && item5 == true
	for i < n-1 && valid {
		fmt.Scan(&item1, &item2, &item3, &item4, &item5)
		i++
		valid = item1 == true && item2 == true && item3 == true && item4 == true && item5 == true
	}
	fmt.Print(valid)
}

/*
Program Pramuka
KD
	n,i : integer
	item1,item2,item3,item4,item5,valid : boolean
Algoritma
	input n
	i <-- 0
	input (item1, item2, item3, item4, item5)
	valid <-- item1 == true and item2 == true and item3 == true and item4 == true and item5 == true
	while i < n-1 and valid do
		input (item1, item2, item3, item4, item5)
		i <-- i + 1
		valid <-- item1 == true and item2 == true and item3 == true and item4 == true and item5 == true
	endwhile
	Print(valid)
*/
