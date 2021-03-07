package main

import "fmt"

func main() {
	meatFactory, err := getFoodFactory("meat")
	if err != nil {
		fmt.Println(err)
	}

	veganFactory, err := getFoodFactory("vegan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", meatFactory.makeBurger())
	fmt.Printf("%+v\n", veganFactory.makeBurger())
}
