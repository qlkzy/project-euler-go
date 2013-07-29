package main

import "fmt"

const maxprimes = 10001

func main() {
	var nprimes = 1
	var primes [maxprimes + 1]int
	primes[0] = 2
	for nprimes <= maxprimes {
		var cand = primes[nprimes-1] + 1

		for di := 0; di < nprimes-1; di++ {
			if cand%primes[di] == 0 {
				di = 0
				cand++
			}
		}

		primes[nprimes] = cand
		nprimes++
	}
	fmt.Println(primes[nprimes-1])
}
