// package notmain

// import "fmt"

// const N = 100

// func odds(in <-chan int, out chan<- int, done chan<- struct{}) {
// 	for i, j := 0, 3; i < N*5; i, j = i+1, j+2 {
// 		select {
// 		case <-in:
// 			break
// 		default:
// 			out <- j
// 		}
// 	}
// }

// func sieve(in <-chan int, out chan<- int) {
// 	prime := <-in
// 	for val := range in {
// 		if val%prime != 0 {
// 			out <- val
// 		}
// 	}
// 	fmt.Println(prime)
// 	close(out)
// }

// func main() {
// 	done := make(chan struct{})
// 	defer close(done)

// 	oCh := make(chan int)

// 	var pCh [N - 1]chan int
// 	for i := range pCh {
// 		pCh[i] = make(chan int)
// 	}
// }
