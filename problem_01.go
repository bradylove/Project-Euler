package main

import (
	"fmt"
)

func main() {
	count := 1000

	for i := 0; i < count; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		} else if i%5 == 0 {
			fmt.Println(i)
		}
	}
}
