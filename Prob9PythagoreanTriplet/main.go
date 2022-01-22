package main

import "fmt"

func main() {
Loop:
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c < 1000; c++ {
				if (a*a+b*b) == c*c && a+b+c == 1000 {
					fmt.Println(a, b, c)
					break Loop
				}
			}
		}
	}
}
