package main

import "fmt"

func main() {
	var result int
	
	for i:=0; i<10; i++ {
		for j:=0; j<10; j++ {
			if i==j {
				continue;
			}
			result = (10*i+j) + (10*j+i)
			if result==99 {
				fmt.Printf("%d%d + %d%d = 99\n", i, j, j, i)
			}
		}
	}
}