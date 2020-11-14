package main

import "fmt"

func main() {
	var (
		card      string
		ofType    = "Spade"
		dealt     = true
		frequency = 0.43
		cards     []string
	)
	card = "Ace it"
	cards = append(cards, card)
	fmt.Println(card, ofType, dealt, frequency)
	fmt.Println(newCard())
	//fmt.Println(cards)
	cards = append(cards, newCard())
	for _, card := range cards {
		fmt.Println(card)
	}
	fmt.Println(cards)
}

func newCard() string {
	// slice can grow and shrink but arrays cannot.
	// arrays and slices are defined with similar datatype
	return "Five of Diamond"
}
