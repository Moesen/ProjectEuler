package main

import "fmt"

const n int = 100

func sumOfSquares(arr [n]int) int {
	sum_square := 0
	for _, val := range arr {
		sum_square += val * val
	}
	return sum_square
}

func squareOfSums(arr [n]int) int {
	square_sum := 0
	for _, val := range arr {
		square_sum += val
	}
	return square_sum * square_sum
}

func main() {
	nats := [n]int{}
	for i := 0; i < n; i++ {
		nats[i] = i + 1
	}

	sumOfSquare := sumOfSquares(nats)
	squareOfSum := squareOfSums(nats)

	fmt.Println(squareOfSum - sumOfSquare)

}
