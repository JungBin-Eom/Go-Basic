package main

import "fmt"

func main() {
	var people []string
	var name string

	for {
		fmt.Scanln(&name)
		if name != "0" {
			people = append(people, name)
		} else {
			break
		}
	}

	for _, person := range people {
		defer fmt.Println(person)
	}
}
