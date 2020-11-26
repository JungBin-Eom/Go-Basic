package main

import "fmt"

func main() {
	names := make([]string, 0, 0)

	var name string
	var result string

	for name != "1" {
		fmt.Scanln(&name)
		names = append(names, name)
		if len(name) > len(result) {
			result = name
		}
	}
	fmt.Println(result, len(result))
}
