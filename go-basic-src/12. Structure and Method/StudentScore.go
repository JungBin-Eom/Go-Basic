package main

import "fmt"

type student struct {
	name   string
	gender string
	score  map[string]int
}

func newStudent() student {
	stud := student{}
	stud.score = map[string]int{}
	return stud
}

func main() {
	var studNum, subNum, score int
	var name, gender, subject string

	var students []student

	fmt.Scanln(&studNum, &subNum)

	for i := 0; i < studNum; i++ {
		person := newStudent()

		fmt.Scanln(&name, &gender)

		person.name = name
		person.gender = gender

		for j := 0; j < subNum; j++ {
			fmt.Scanln(&subject, &score)
			person.score[subject] = score
		}

		students = append(students, person)
	}

	for _, person := range students {
		fmt.Println("----------")
		fmt.Println(person.name, person.gender)

		for index, val := range person.score {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
