package main

import (
	"errors"
	"fmt"
)

func splitStringAndReverse(s string) (string, string, error) {
	if len(s)%2 != 0 {
		return "", "", errors.New("This no work")
	}

	f, temp := s[:len(s)/2], s[len(s)/2:]

	b := ""
	for _, v := range temp {
		b = string(v) + b
	}

	return f, b, nil
}

func main() {
	max := 0
	for i := 999; i > 0; i-- {
		for j := 999; j > 0; j-- {
			f, b, err := splitStringAndReverse(fmt.Sprintf("%d", (i * j)))
			if err == nil && f == b {
				if i*j > max {
					max = i * j
				}
			}
		}
	}
	fmt.Println(max)
}
