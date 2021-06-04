package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer close(c)
		c <- 111
		c <- 222
	}()

	i, ok := <-c
	fmt.Println(i, ok)

	i, ok = <-c
	fmt.Println(i, ok)

	i, ok = <-c
	fmt.Println(i, ok)
}
