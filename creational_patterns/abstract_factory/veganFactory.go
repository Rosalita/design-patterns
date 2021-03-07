package main

type veganFactory struct{}

func (v *veganFactory) makeBurger() burgeriface {
	return &veganBurger{
		burger: burger{
			calories:     177,
			containsMeat: false,
		},
	}
}
