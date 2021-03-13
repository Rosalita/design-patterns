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

//ButtonBuilder holds the data used to construct a button.
type ButtonBuilder struct {
	DisplayText      string
	TextColour       string
	BackgroundColour string
}

//ButtonBuilderIface is the interface representing the chain of methods that build a button.
type ButtonBuilderIface interface {
	DisplayText(string) ButtonBuilderIface
	TextColour(string) ButtonBuilderIface
	BackgroundColour(string) ButtonBuilderIface
	Build() Button
}

//NewButtonBuilder instantiates a new ButtonBuilder
func NewButtonBuilder() ButtonBuilder {
	return ButtonBuilder{}
}

//SetDisplayText sets the display text on the button builder.
//Then passes the button builder to the next method in the chain.
func (bb *ButtonBuilder) SetDisplayText(text string) *ButtonBuilder {
	bb.DisplayText = text
	return bb
}

//SetTextColour sets the text colour on the button builder.
//Then passes the button builder to the next method in the chain.
func (bb *ButtonBuilder) SetTextColour(colour string) *ButtonBuilder {
	bb.TextColour = colour
	return bb
}

//SetBackgroundColour sets the text colour on the button builder.
//Then passes the button builder to the next method in the chain.
func (bb *ButtonBuilder) SetBackgroundColour(colour string) *ButtonBuilder {
	bb.BackgroundColour = colour
	return bb
}

//Build generates the button from all the data which has been set on the button builder.
func (bb *ButtonBuilder) Build() Button {
	return Button{
		displayText:      bb.DisplayText,
		textColour:       bb.TextColour,
		backgroundColour: bb.BackgroundColour,
	}
}

func main() {

	ButtonBuilder := NewButtonBuilder()

	myButton := ButtonBuilder.SetDisplayText("Hello World").SetTextColour("blue").SetBackgroundColour("white").Build()

	// This method chaining is often split over multiple lines, e.g.
	anotherButton := ButtonBuilder.SetDisplayText("foo").
		SetTextColour("bar").
		SetBackgroundColour("baz").
		Build()

	fmt.Printf("%+v\n", myButton)
	fmt.Printf("%+v\n", anotherButton)
}
