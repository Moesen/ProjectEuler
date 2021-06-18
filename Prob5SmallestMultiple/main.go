package main

import (
	"fmt"
	"math/big"
)

// Cake-/Ladder-Method used
// https://www.calculatorsoup.com/calculators/math/lcm.php

func primeGenerator() func() int {
	x := 1
	return func() int {
		for {
			x++
			if big.NewInt(int64(x)).ProbablyPrime(0) {
				return x
			}
		}
	}
}

func main() {
	const num int = 20

	arr := [num]int{}
	for i := 0; i < num; i++ {
		arr[i] = i + 1
	}

	generatePrime := primeGenerator()
	primes := []int{}

	for {
		prime := generatePrime()
		for {
			changed := false
			for i, val := range arr {
				if val%prime == 0 && val != 1 {
					arr[i] = val / prime
					changed = true
				}
			}
			if changed {
				primes = append(primes, prime)
			} else {
				break
			}
		}
		sum := 0
		for _, val := range arr {
			sum += val
		}
		if sum == len(arr) {
			break
		}
	}

	fmt.Println(primes)
	product := 1
	for _, prime := range primes {
		product *= prime
	}

	fmt.Println(product)
}
