//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
	fmt.Println(start, time.Now())
}
