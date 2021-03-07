package main

type veganBurger struct {
	burger
}

func (v *veganBurger) setCals(calories int) {
	v.calories = calories
}

func (v *veganBurger) getCals(calories int) int {
	return v.calories
}
