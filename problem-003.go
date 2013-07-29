package main

import "fmt"

func main() {
	var num, factor uint64 = 600851475143, 1
	for num > 1 {
		factor++
		if num%factor == 0 {
			num /= factor
			if num > 1 {
				factor = 1
			}
		}
	}
	fmt.Println(factor)
}
