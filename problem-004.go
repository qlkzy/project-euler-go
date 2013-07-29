package main

import "fmt"

func digits(x int) []int {
	ds := []int{}
	for x > 0 {
		ds = append(ds, x%10)
		x /= 10
	}
	return ds
}

func palindrome(x int) bool {
	ds := digits(x)
	var f, b = 0, len(ds) - 1
	if b < f {
		return false
	}
	for f < b {
		if ds[f] != ds[b] {
			return false
		}
		f++
		b--
	}
	return true
}

func main() {
	var largest = 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			var product = x * y
			if palindrome(product) && product > largest {
				largest = product
			}
		}
	}
	fmt.Println(largest)
}
