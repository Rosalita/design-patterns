package main

import (
	"errors"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) getName() string {
	return p.Name
}

func (p Person) getAge() int {
	return p.Age
}

// This is an example of the simplest factory in Go
// A function that takes some parameters and creates a struct.
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// A Pointer to the created object can also be returned.
// This is handy if an error occurs as no object can be returned.
func NewPerson2(name string, age int) (*Person, error) {
	if age < 0 {
		return nil, errors.New("person age must be a positive integer")
	}
	return &Person{
		Name: name,
		Age:  age,
	}, nil
}
