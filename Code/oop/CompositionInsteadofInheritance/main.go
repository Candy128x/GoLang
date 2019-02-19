package main


import (
	"fmt"
	)


type author struct {
	firstName, lastName, bio string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}


type post struct {
	title, content string
	author
}

func (p post) details() {
	fmt.Println("Title:", p.title)
	fmt.Println("Content:", p.content)
	fmt.Println("Author:", p.fullName())
	fmt.Println("Bio:", p.bio)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    author1 := author {
    	"Shamu",
    	"Dave",
    	"Writer",
    }

    post1 := post {
    	"In haritance in Go..",
    	"Go supports composition instead of inheritance",
    	author1,
    }
    
    post1.details()

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/oop/CompositionInsteadofInheritance$ go run main.go 
Hello Go Developer.. :]
Title: In haritance in Go..
Content: Go supports composition instead of inheritance
Author: Shamu Dave
Bio: Writer
*/


/* Title
-> Description..
*/
