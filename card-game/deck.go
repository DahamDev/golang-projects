package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	names := []string{"Diamonds","Hearts","Spades","Clubs"}
	values :=[]string{"Ace","One","Two","Three","Four"}

	new_deck := deck{}
	for _, name := range names{
		for _, value := range values{
			new_deck = append(new_deck, name+"-"+value)
		}
	}

	return new_deck
}

//this is call a receiver function. (d deck) will enable function to be called as deck.print() instead of parsing 
//deck object as an parameter to the function.
func (d deck) print(){
	fmt.Println(d)
}

func (d deck) dual(size int) (deck,deck){
	return d[:size],d[size:]
}

func (d deck) writeToFile(fileName string){
	deckToString  := strings.Join(d, ",")
	byteArray := []byte(deckToString)
	err := os.WriteFile(fileName, byteArray, 0644)
	if err != nil {
		log.Fatal(err)
	}
	
}

func getDeckFromFile(fileName string) (deck){
	b, err := os.ReadFile(fileName)
	if err != nil{
		log.Fatal("Error reading from file "+err.Error())
	}
	stringArray := strings.Split(string(b),",")
	return deck(stringArray)
}

func (d deck) shuffle(){
	
	for index,_ := range d{
		position := rand.Intn(len(d)-1)
		d[index], d[position] = d[position], d[index]
	}
}