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
	
    // assign multiple val to multiple vari
    var i, s = 2, "hi" // i = 2 & s = "hi"
	fmt.Println(i, s)
	// op: 2 hi
	
	var length, width, height int = 10, 20, 30 // Declaring multiple variables 
	fmt.Println("Length is:", length, " Width is:", width, " Height is:", height)
	
	var lengthNew, widthNew, heightNew = 10, 20, 30 // "int" is ommited/dropped 
	fmt.Println("New.. Length is:", lengthNew, " Width is:", widthNew, " Height is:", heightNew)

	
	// Declare variables belonging to different types in a single statement.
	var (
		subject = "Python"
		percent	= 85.3
		rank	= 1
	)
	fmt.Println("Hey.. In", subject, "subject i got", percent, "percent with", rank, "rank :)")
	
	
	// Short hand declaration
	fruit, price := "Apple", 62
	fmt.Println("Furit:", fruit, "& Price:", price)
	/*
	Short hand syntax can only be used when at least one of the variables in the left side of := is 
		newly declared.
	*/


	pi := float32 (21)
	fmt.Printf("value of pi: %f & type: %T", pi, pi)
	fmt.Printf("\nvalue of pi: %.2f & type: %T", pi, pi)


	fmt.Println("\n\n") // The End.. .
}

/* 
? What is a variable
-> Variable is the name given to a memory location to store a value of a specific type.
*/