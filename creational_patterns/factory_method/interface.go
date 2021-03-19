package main

import (
	"errors"
)

type Personiface interface {
	getName() string
	getAge() int
}

// Factory functions can also return interfaces.
func NewPerson3(name string, age int) Personiface {
	return Person{
		Name: name,
		Age:  age,
	}
}

// A Pointer to a created object can pass through a interface it satisfies.
func NewPerson4(name string, age int) (Personiface, error) {
	if age < 0 {
		return nil, errors.New("person age must be a positive integer")
	}
	return &Person{
		Name: name,
		Age:  age,
	}, nil
}

// Accept interfaces return structs
func NewPerson5(p Personiface) Person {
	return Person{
		Name: p.getName(),
		Age:  p.getAge(),
	}
}
