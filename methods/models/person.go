package models

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

// Presentation - Greet everyone and show up
func (p Person) Presentation() string {
	return fmt.Sprintf("Hi, everyone! I'm %s %s and I have %d years old.\n", p.Name, p.LastName, p.Age)
}
