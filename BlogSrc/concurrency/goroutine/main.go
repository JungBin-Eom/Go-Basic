package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Ricky", " - ", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Tom", " - ", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("Alice", " - ", i)
		}
	}()

	wg.Wait()
}
