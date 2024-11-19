package main

import (
	"fmt"
	dto "nutshell/methods/models"
)

func main() {
	student := dto.Student{
		Person: dto.Person{
			Name:     "Joe",
			LastName: "Doe",
			Age:      33,
		},
		Class: "5A",
	}
	showUp(student)
}

func showUp(p dto.Presenter) {
	fmt.Println(p.Presentation())
}
