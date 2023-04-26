package main

import "fmt"

func main() {
	var n, i int
	var item1, item2, item3, item4, item5 bool
	var ready bool
	fmt.Scan(&n)
	fmt.Scan(&item1, &item2, &item3, &item4, &item5)
	ready = item1 && item2 && item3 && item4 && item5
	i = 0
	for i < n-1 && ready {
		fmt.Scan(&item1, &item2, &item3, &item4, &item5)
		ready = item1 && item2 && item3 && item4 && item5
		i++
	}
	fmt.Print(ready)
}
