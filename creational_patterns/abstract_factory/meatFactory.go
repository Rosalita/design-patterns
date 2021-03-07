package main

type meatFactory struct{}

func (m *meatFactory) makeBurger() burgeriface {
	return &meatBurger{
		burger: burger{
			calories:     204,
			containsMeat: true,
		},
	}
}
