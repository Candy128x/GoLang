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


type website struct {
	posts []post // we have added the field name posts to the slice of post []post.
}

func (w website) contents() {
	fmt.Println("Content of Website: \n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
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
    
    post2 := post {
    	"Struct instead of class in go..",
    	"Go does not support classes but method can be added to structs",
    	author1,
    }

    post3 := post {
    	"Concurrency",
    	"Go is a concurrent lang not a parallel one",
    	author1,
    }

    w := website {
    	posts: []post { post1, post2, post3 },
    }
    w.contents()

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/oop/CompositionInsteadofInheritance$ go run EmbeddingSliceofStructs.go 
Hello Go Developer.. :]
Content of Website: 

Title: In haritance in Go..
Content: Go supports composition instead of inheritance
Author: Shamu Dave
Bio: Writer

Title: Struct instead of class in go..
Content: Go does not support classes but method can be added to structs
Author: Shamu Dave
Bio: Writer

Title: Concurrency
Content: Go is a concurrent lang not a parallel one
Author: Shamu Dave
Bio: Writer

*/


/* Embedding slice of structs
-> Description..
*/