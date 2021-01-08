package Concurrency_Pattern

import "fmt"

// PlusOne returns a channel of num+1 for received from in.
func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}
