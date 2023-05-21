package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){

	personOne  := person {
		name: "peter",
		age: 25,
	}
	personOne.updateName("david")
	//if we print the person struc it wont get updated as navinda. updateName function will create a new 
	// struct and poin the output of the function to the new pointer.
	fmt.Printf("after updating without pointer = %+v \n", personOne)

	personOnePointer := &personOne //get the memory address of the personOne
	personOnePointer.updateNamePointer("John") //instead of parsing the pointer here we can directly pass the object as well

	fmt.Printf("after updating with pointer = %+v", personOne)
}

//what will happen when you call this function?
//reseiver function will create a new person struct and point the result of this function
//to the new struct. it will not udpate the original struct
func (p person) updateName(name string){
	p.name = name
}

//*person means pointer to the person object
// *pointerToPerson means get the actual value at the pointer

func (pointerToPerson *person) updateNamePointer(name string){
	(*pointerToPerson).name = name
}


//go has two different types of data. reference types and value types
// values types are : int, float, boolean, string, arrays
// reference types are : slice

// when we create a slice, go will automatically create two data structure
// 1.array
// 2.another set of poitner to point to the original array and size

// when we pass a slice to an method, method will create a copy of the slice data structure, 
// but it willbe pointing to the same origina array object
// Therefor, when we dont have worry about pointers when passing slices to methods