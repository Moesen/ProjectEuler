package main

import "fmt"

const N = 10001

func odds(primeCh chan<- int, in <-chan int, out chan<- int) {
Loop:
	for i, j := 0, 3; i < N*100; i, j = i+1, j+2 {
		select {
		case <-in:
			break Loop
		default:
			out <- j
		}

	}

	primeCh <- 2
	close(out)
	for {
		if _, ok := <-in; !ok {
			close(primeCh)
			break
		}
	}
}

func sieve(primeCh chan<- int, in <-chan int, out chan<- int) {
	prime := <-in
	for val := range in {
		if val%prime != 0 {
			out <- val
		}
	}

	primeCh <- prime
	close(out)
}

func generatePrimes(numOfPrimes int) []int {

	ch1 := make(chan int)
	var chs []chan int
	for i := 0; i < numOfPrimes; i++ {
		chs = append(chs, make(chan int))
	}

	primeCh := make(chan int)

	go odds(primeCh, chs[numOfPrimes-2], ch1)
	go sieve(primeCh, ch1, chs[0])
	for i := 0; i < numOfPrimes-2; i++ {
		go sieve(primeCh, chs[i], chs[i+1])
	}

	primes := []int{}
	for prime := range primeCh {
		primes = append(primes, prime)
	}

	return primes
}

func main() {
	fmt.Println(generatePrimes(5))
}
