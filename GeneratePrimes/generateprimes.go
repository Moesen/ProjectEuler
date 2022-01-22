package GeneratePrimes

import (
	"math"
)

func odds(primeCh chan<- int, in <-chan int, out chan<- int, n int) {
Loop:
	for i, j := 0, 3; i < n*100; i, j = i+1, j+2 {
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

func generatePrimesN(numOfPrimes int) []int {
	const N = 10001

	ch1 := make(chan int)
	var chs []chan int
	for i := 0; i < numOfPrimes; i++ {
		chs = append(chs, make(chan int))
	}

	primeCh := make(chan int)

	go odds(primeCh, chs[numOfPrimes-2], ch1, N)
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

func generatePrimesLim(lim int) []int {
	estimate := int(lim / int(math.Log(float64(lim))))
	estimate = int(float64(estimate) * 1.5)

	ch1 := make(chan int)
	var chs []chan int
	for i := 0; i < estimate; i++ {
		chs = append(chs, make(chan int))
	}

	primeCh := make(chan int)

	go odds(primeCh, chs[estimate-2], ch1, estimate)
	go sieve(primeCh, ch1, chs[0])
	for i := 0; i < estimate-2; i++ {
		go sieve(primeCh, chs[i], chs[i+1])
	}

	primes := []int{}
	for prime := range primeCh {
		if prime <= lim {
			primes = append(primes, prime)
		} else {
			break
		}
	}

	return primes
}
