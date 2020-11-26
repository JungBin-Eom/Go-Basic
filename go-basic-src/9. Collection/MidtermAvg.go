package main

import "fmt"

func main() {
	var midterm = make(map[string]int)
	var avg float32
	var subject string
	var score, sum int

	fmt.Scanln(&subject, &score)

	for subject != "0" {
		midterm[subject] = score
		sum += score
		fmt.Scanln(&subject, &score)
	}

	avg = float32(sum) / float32(len(midterm))

	for index, num := range midterm {
		fmt.Printf("%s %d\n", index, num)
	}

	fmt.Printf("%.2f", avg)
}
