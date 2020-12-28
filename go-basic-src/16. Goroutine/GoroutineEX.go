package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a int, b int, result *int, w *sync.WaitGroup) {
	defer w.Done()
	*result = a + b
}

func main() {
	var num1, num2 int = 10, 5
	var result int
	wait := new(sync.WaitGroup)
	wait.Add(1)

	go add(num1, num2, &result, wait)

	time.Sleep(time.Second)
	fmt.Println(result)

	wait.Wait()
}
