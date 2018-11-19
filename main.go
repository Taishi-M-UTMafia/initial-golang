package main

import "fmt"

// ---------------------------------------------
// Cards
// ---------------------------------------------
// func main() {
// cards := newDeck()

// // receiver on function
// // cards.print()

// hand, remainingCards := deal(cards, 5)

// hand.print()
// remainingCards.print()

// cards.saveToFile("MyCards")

// cards := newDeckFromFile("MyCards")
// cards.print()

// cards := newDeck()
// cards.shuffle()
// cards.print()
// }

// ---------------------------------------------
// Structs
// ---------------------------------------------
// func main() {
// alex := person{firstName: "Alex", lastName: "Anderson"}
// var alex person

// alex.firstName = "Alex"
// alex.lastName = "Anderson"

// fmt.Println(alex)
// fmt.Printf("%+v", alex)

// jim := person{
// 	firstName: "Jim",
// 	lastName:  "Party",
// 	contactInfo: contactInfo{
// 		email:   "jimpgmail.com",
// 		zipCode: 94000,
// 	},
// }

// jimPointer := &jim
// jimPointer.updateName("Jimmy")

// updateNameのreceiverのtypeは*personな一方で
// jimの方はただのpersonだが、そこはgoがよしなにやってくれる
// jim.updateName("Jimmy")
// jim.print()
// }

// ---------------------------------------------
// Maps
// ---------------------------------------------
func main() {
	// structと似てる！

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4327as",
	// }

	// var colors map[string]string

	colors := make(map[string]string)
	// structだとcolors.white = hogeみたいな感じになるという相違点
	// keyとvalueの肩が指定されている(structはされてない)
	colors["white"] = "#fff"

	delete(colors, "white")

	fmt.Println(colors)
}
