package main

import "fmt"

//Button is a UI element that displays text, it has a text colour and a background colour.
//It has getter methods to expose the data it contains.
type Button struct {
	displayText      string
	textColour       string
	backgroundColour string
}

func (b *Button) getDisplayText() string {
	return b.displayText
}

func (b *Button) getTextColour() string {
	return b.textColour
}

func (b *Button) getBackgroundColour() string {
	return b.backgroundColour
}

//ButtonInput is used to create a button.
type ButtonInput struct {
	DisplayText      string
	TextColour       string
	BackgroundColour string
}

//NewButton creates a new button.
func NewButton(input ButtonInput) Button {
	return Button{
		displayText:      input.DisplayText,
		textColour:       input.TextColour,
		backgroundColour: input.BackgroundColour,
	}
}

func main() {

	//myButton := ButtonBuilder.SetDisplayText("Hello World").SetTextColour("blue").SetBackgroundColour("white").Build()
	myButton := NewButton(ButtonInput{DisplayText: "Hello World", TextColour: "blue", BackgroundColour: "white"})

	// This instantiation can also be split over muliple lines
	anotherButton := NewButton(ButtonInput{
		DisplayText:      "foo",
		TextColour:       "bar",
		BackgroundColour: "baz",
	})

	fmt.Printf("%+v\n", myButton)
	fmt.Printf("%+v\n", anotherButton)
}
