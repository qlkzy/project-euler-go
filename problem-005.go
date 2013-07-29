package main

import "fmt"

func divisible(x int) bool {
	for d := 1; d <= 20; d++ {
		if x%d != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 1; ; i++ {
		if divisible(i) {
			fmt.Println(i)
			return
		}
	}
}
