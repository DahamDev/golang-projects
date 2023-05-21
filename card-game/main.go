package main

import "fmt"

func main() {

	deck := newDeck()
	fmt.Print("new deck is ")
	fmt.Println(deck)

	firstdual, secondual := deck.dual(5)
	firstdual.print()
	secondual.print()

	firstDual := getDeckFromFile("file1.txt")
	print("Before shuffle ")
	firstDual.print()

	print("After shuffle")
	firstDual.shuffle()
	firstDual.print()
}