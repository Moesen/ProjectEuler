package main

import (
	"fmt"
)

const N = 10001

func odds(in <-chan uint64, out chan<- uint64, doneCh chan<- struct{}) {
Loop:
	for i, j := 0, 3; i < N*100; i, j = i+1, j+2 {
		select {
		case <-in:
			break Loop
		default:
			send := uint64(j)
			out <- send
		}

	}
	fmt.Println(2)
	close(out)

	for {
		if _, ok := <-in; !ok {
			doneCh <- struct{}{}
			break
		}
	}
}

func sieve(in <-chan uint64, out chan<- uint64) {
	prime := <-in
	for val := range in {
		if val%prime != 0 {
			out <- val
		}
	}

	fmt.Println(prime)
	close(out)
}

func main() {
	doneCh := make(chan struct{})
	defer close(doneCh) // Seemed like a good place to instruct it to close channel on main end

	oCh := make(chan uint64)

	var pCh [N - 1]chan uint64
	for i := range pCh {
		pCh[i] = make(chan uint64)
	}

	fmt.Println("The first", N, "prime numbers are:")

	// Connect/start goroutines
	go odds(pCh[N-2], oCh, doneCh)
	go sieve(oCh, pCh[0])
	for i := 0; i < N-2; i++ {
		go sieve(pCh[i], pCh[i+1])
	}

	// Await termination
	<-doneCh

	fmt.Println("Done!")
}
