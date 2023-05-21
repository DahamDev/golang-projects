package main

import "fmt"

type contactInfo struct {
	email string
	address string
}

type person struct {
	firstName string
	secondName string
	contactDetails contactInfo
}

func main(){

	people := []person{
		{firstName:"daham", secondName: "navinda", contactDetails: contactInfo{email:"navindabc@gmail.com", address:"madelgamuwa"}},
		{firstName: "faith", secondName: "culas"}}

	for i, person := range people{
		fmt.Printf("%d person is %s . details : %+v",i, person.firstName, person)
	}

}