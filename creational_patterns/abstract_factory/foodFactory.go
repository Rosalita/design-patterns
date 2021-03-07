package main

import (
	"errors"
)

type foodFactory interface {
	makeBurger() burgeriface
}

func getFoodFactory(diet string) (foodFactory, error) {
	if diet == "meat" {
		return &meatFactory{}, nil
	}

	if diet == "vegan" {
		return &veganFactory{}, nil
	}
	return nil, errors.New("unknown diet")
}
