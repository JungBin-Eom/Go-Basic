package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 24
}

func main() {
	user1 := User{Name: "ricky", Email: "ricky@gmail.com", Age: 25}
	user2 := User{Name: "venny", Email: "venny@gmail.com", Age: 23}
	users := []User{user1, user2}
	// {{}}부분이 비어있는 부분
	// 비어있는 부분을 채우자
	tmpl, err := template.New("Teml1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
