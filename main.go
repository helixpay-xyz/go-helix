package main

import (
	"fmt"
	"sync"
	"time"
)

func print(number uint) {
	// Print the number to the console
	fmt.Println(number)
}

func main() {
	start := time.Now() // Start timer

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(n uint) {
			defer wg.Done()
			print(n)
		}(uint(i))
	}

	wg.Wait()

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Execution time: %s\n", elapsed)
}
