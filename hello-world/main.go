package main

import (
	"fmt"
	"strconv"
)

//function to check odd and even numbers
func main(){
	//define a struct with numbers rom 0, 10
	numbers:=[]int{0,1,2,3,4,5,6,7,8,9,10}
	for _, number := range numbers{
		if number%2 == 0{
			fmt.Println(strconv.Itoa(number) + " is an even number")
		}
	}
}

