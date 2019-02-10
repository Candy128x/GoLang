package main

import "fmt"

func main(){
	var age int // Variable declaration
	fmt.Println("Your age is:", age) // initial value is Zero. Output: Your age is: 0

	var name = "Ashish"
	fmt.Println("Your name: " + name)
	
	age = 12 // Assignment
	fmt.Println("Your age is:", age)
	age = 23 // Assignment, reuse variable
	fmt.Println("Your new age is:", age)
	
	var ageVote int = 18 // Variable declare with initial value
	fmt.Println("Your Vote age is:", ageVote)
	
	var ageMovie = 21 // type will be inferred
	fmt.Println("Your are eligible to watch movie when your age is:", ageMovie)
	/*
	If a variable has an initial value, Go will automatically be able to infer the type of that 
		variable using that initial value. Hence if a variable has an initial value, the type in 
		the variable declaration can be omitted.
	If the variable is declared using the syntax var name = initialvalue, Go will automatically 
		infer the type of that variable from the initial value.
	*/
	
}

/* 
? What is a variable
-> Variable is the name given to a memory location to store a value of a specific type.
*/