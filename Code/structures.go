package main


import (
	"fmt"
	)


	/* The above Employee struct is called a named structure because it creates a new type named Employee 
		which can be used to create structures of type Employee.
	*/
    type Employee struct { // named structure
    	firstName string
    	age, salary int
    }

    /* It is possible to declare structures without declaring a new type and these type of structures are 
    called anonymous structures.
    */
    var employee struct { // anonymous structure
    	firstName string
    	age, salary int
    }


    // Anonymous fields
    type Person struct {
    	string
    	int 
    }


    // Nested structs
    type Address struct {  
    	city, state string
	}
	type PersonNested struct {  
	    name string
	    age int
	    address Address
	}


	// Promoted fields
	type AddressPromoted struct {  
    	city, state string
	}
	type PersonPromoted struct {  
	    name string
	    age  int
	    AddressPromoted
	}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    fmt.Println("Values of employee:", employee)

    emp1 := Employee {
    	firstName : "Raju",
    	age : 23,
    	salary : 8000,
    }
    fmt.Println("Values of emp1:", emp1)


    emp2 := Employee{"Shamu", 21, 9500}
    fmt.Println("Values of emp2:",emp2)



    // Creating anonymous structures 
    car := struct {
    	company, model string
    	price int
    }{
    	company : "Audi",
    	model : "class-C",
    	price : 2,
    }
    fmt.Println("Values of car:", car)
    fmt.Println("Values of car -> Audi -> price:", car.price)


    // Anonymous fields
    p := Person{"Modi", 50}
    fmt.Println("p", p)

    var p1 Person
    p1.string = "Gandhi"
    p1.string = "Gandhi2" // overwrite
    p1.int = 23
    fmt.Println("p1:", p1)


    // Nested structs
    fmt.Println("\n--Nested structs--")
    var pNested PersonNested
    pNested.name = "Naveen"
    pNested.age = 50
    pNested.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", pNested.name)
    fmt.Println("Age:",pNested.age)
    fmt.Println("City:",pNested.address.city)
    fmt.Println("State:",pNested.address.state)


    // Promoted fields
    fmt.Println("\n--Promoted fields--")
    var pPromoted PersonPromoted
    pPromoted.name = "Naveen"
    pPromoted.age = 50
    pPromoted.AddressPromoted = AddressPromoted{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", pPromoted.name)
    fmt.Println("Age:", pPromoted.age)
    fmt.Println("City:", pPromoted.city) //city is promoted field
    fmt.Println("State:", pPromoted.state) //state is promoted field


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run structures.go 
Hello Go Developer.. :]
Values of employee: { 0 0}
Values of emp1: {Raju 23 8000}
Values of emp2: {Shamu 21 9500}
Values of car: {Audi class-C 2}
Values of car -> Audi -> price: 2
p {Modi 50}
p1: {Gandhi2 23}

--Nested structs--
Name: Naveen
Age: 50
City: Chicago
State: Illinois

--Promoted fields--
Name: Naveen
Age: 50
City: Chicago
State: Illinois
*/

/*
What is a structure?
A structure is a user defined type which represents a collection of fields. 
*/