package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
