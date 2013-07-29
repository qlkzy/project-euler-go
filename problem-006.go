package main

import "fmt"

func main() {
	var sumsq, sum = 0, 0
	for i := 1; i <= 100; i++ {
		sumsq += i * i
		sum += i
	}
	diff := sum*sum - sumsq
	fmt.Println(diff)
}
