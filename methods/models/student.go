package models

import "fmt"

type Student struct {
	Person
	Class string
}

// Presentation - Greet the teacher and show up
func (s Student) Presentation() string {
	return fmt.Sprintf("Good morning teacher, I'm %s %s.\n", s.Person.Name, s.Person.LastName)
}
