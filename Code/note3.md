# `Go Lang Notes </>`


---
## Comments
1. "//" for Single line
2. "/*" 
	for 
	multiple
	line
  "*/"


---
## Append vari in print statment

1. use '+' for string
> Eg: var name = "Ashish"
~ fmt.Println("Your name: " + name)

2. use ',' for string, int 
- It's automatically not append one space in prefix of variable
> Eg: var ageTeen = 18
~ fmt.Println("Teen age:", ageTeen)


---
## method() name declaration
1. smallAlphbetFunc() to acceess with in single own file like private
2. CapsFunc() to access globely or public


---
## use of %T 
- The type of a variable can be printed using %T format specifier in Printf method.
> Eg: 
a := 23
fmt.Printf("type of a is %T", a) //type a	
~ - Output: type of a is: %T 23

## use of %d
- Go has a package unsafe which has a Sizeof function which returns in bytes the size of the variable 
	passed to it
> Eg: 
a := 23
fmt.Printf("size of a is %d", unsafe.Sizeof(a)) // size of a	
~ - Output: size of a is: %d 8


---
## use of defer


---
## panic






---
Author: [Ashish](https://github.com/Candy128x) </>
