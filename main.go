// This sample program demonstrates how to create goroutines and
// how the goroutine scheduler behaves with three logical processors.

package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	// Allocate three logical processors for the scheduler to use.
	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup
	names := []string{"Marín", "Berny", "Tony"}

	//Add a count depends on the length of names array
	wg.Add(len(names))

	for _, name := range names {
		//Anonymous function.
		go func(name string) {
			//defers the execution of a function until
			//the surrounding function returns.
			defer wg.Done()
			r := 0.0
			for i := 0; i < 100000000; i++ {
				r += math.Pi * math.Sin(float64(len(name)))
			}
			fmt.Println("Hello: ", name)
		}(name)

	}

	// Wait for the goroutines to finish.
	wg.Wait()
}
