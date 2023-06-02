package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

type bot interface {
	getGreeting() string //all the structs which have getGreeting method which return a strung
						 // will become type bot``
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	english_greeting := eb.getGreeting()
	spanish_greeting := sb.getGreeting()

	fmt.Printf("english greeting %s  , spanisg greeting %s \n", english_greeting, spanish_greeting)

	//printing using the interface function
	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	return "Hi How are you"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!!!!!"
}


func printGreeting( b bot){
	fmt.Printf("Greeting is %s \n", b.getGreeting())
}