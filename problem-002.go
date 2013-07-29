package main

import "fmt"

func main() {
	var curr, prev, total uint64 = 1, 1, 0
	for curr < 4000000 {
		if curr%2 == 0 {
			total += curr
		}
		curr, prev = curr+prev, curr
	}
	fmt.Println(total)
}
