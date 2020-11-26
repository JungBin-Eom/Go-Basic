package main

import "fmt"

func main() {
	var A = [2][2]int{
		{7, 3},
		{5, 2},
	}

	var have bool

	d := A[0][0]*A[1][1] - A[0][1]*A[1][0]

	if d != 0 {
		var Ad = [2][2]int{
			{A[1][1] * d, -A[0][1] * d},
			{-A[1][0] * d, A[0][0] * d},
		}

		have = true

		fmt.Println(have)
		fmt.Println(Ad[0][0], Ad[0][1])
		fmt.Println(Ad[1][0], Ad[1][1])

	} else {
		have = false

		fmt.Println(have)
	}
}
