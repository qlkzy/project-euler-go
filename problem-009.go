package main

import "fmt"

func main() {
	res := 0
	for a := 0; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			for c := b; c < 1000; c++ {
				if (a+b+c == 1000) && (a*a+b*b == c*c) {
					res = a * b * c
					break
				}
			}
		}
	}
	fmt.Println(res)
}
