package main

type meatBurger struct {
	burger
}

func (m *meatBurger) setCals(calories int) {
	m.calories = calories
}

func (m *meatBurger) getCals(calories int) int {
	return m.calories
}
