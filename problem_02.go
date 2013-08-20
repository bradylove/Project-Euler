package main

import (
	"fmt"
)

func main() {
	total := 0

	last := 1
	current := 2

	for {
        if current > 4000000 {
            break
        }

		if current%2 == 0 {
			total += current
		}

		last, current = current, current+last
	}

    fmt.Println("Total:", total)
}
