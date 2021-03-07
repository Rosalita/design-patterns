package main

type burgeriface interface {
	setCals(calories int)
	getCals(calories int) int
}

type burger struct {
	calories     int
	containsMeat bool
}

func (b *burger) setCals(calories int) {
	b.calories = calories
}

func (b *burger) getCals(calories int) int {
	return b.calories
}

func (b *burger) isVegan() bool {
	return !b.containsMeat
}
