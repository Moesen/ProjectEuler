package main

import (
	"fmt"
	"math"
)

const N = 600851475143

func main() {
	n := N

	for n%2 == 0 {
		fmt.Println(2)
		n = n / 2
	}

	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
		for n%i == 0 {
			fmt.Println(i)
			n /= i
		}
	}

	if n > 2 {
		fmt.Println(n)
	}

}
