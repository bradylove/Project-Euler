package main

import (
	"fmt"
)

func main() {
	count := 1000
    total := 0

	for i := 0; i < count; i++ {
		if i%3 == 0 {
			total += i
		} else if i%5 == 0 {
			total += i
		}
	}

    fmt.Println("Total:", total)
}
