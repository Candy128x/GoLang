package main

import (
	"fmt" 
	"unsafe"
)

func main() {

	male := true // boolean
	if male {
	fmt.Println("You are male..")
	}
	// op: You are male..

	a := 23
	fmt.Println("type of a is: %T",a)
	fmt.Println("size of a is: %d", unsafe.Sizeof(a))
	/* Output:
	type of a is: %T 23
	size of a is: %d 8
	*/
	
	// Complex types
	c1 := complex(5, 7)
	c2 := 8 + 27i
	fmt.Println("\n value of c1: & c2:")
	
	cAdd := c1 + c2
	fmt.Println("Complex Addition:",cAdd)
	
	cMin := c1 - c2
	fmt.Println("Complex Minus:",cMin)
	
	cMul := c1 * c2
	fmt.Println("Complex Multiple:",cMul)
	
	/* Compilation logic
	Example: (3 + 2i)(1 + 7i)
	(3 + 2i)(1 + 7i) = 3×1 + 3×7i + 2i×1+ 2i×7i
	= 3 + 21i + 2i + 14i2
	= 3 + 21i + 2i − 14   (because i2 = −1)
	= −11 + 23i
	*/
	
	cDiv := c1 / c2
	fmt.Println("Complex Division:",cDiv)
	
	/* Output:
	Complex Addition: (13+34i)
	Complex Minus: (-3-20i)
	Complex Multiple: (-149+191i)
	Complex Division: (0.28877679697351827-0.09962168978562422i)
	*/
	
	
	// ? Samaj main aaya kya!
	var i, s = 2, "hi"
	fmt.Println(i, s)
	// op: 2 hi
	
	
	first	:= "Ashish"
    last	:= "Developer"
    name	:= first +" "+ last
    fmt.Println("My name is",name ,"i m not terrorist :)")
	// op: My name is Ashish Developer i m not terrorist :)
	
	
	// type Conversion
	i, j := 12, 23.5
	k := i + int(j)
	fmt.Println("\n Value of k:", k)
	// op: Value of k: 35
	
}

/* Output:
You are male..
type of a is: %T 23
size of a is: %d 8
*/


/* Types of type
-> The following are the basic types available in go
- bool
- Numeric Types
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte
	rune
- string


-> Signed integers

int8: represents 8 bit signed integers
size: 8 bits
range: -128 to 127

int16: represents 16 bit signed integers
size: 16 bits
range: -32768 to 32767

int32: represents 32 bit signed integers
size: 32 bits
range: -2147483648 to 2147483647

int64: represents 64 bit signed integers
size: 64 bits
range: -9223372036854775808 to 9223372036854775807

int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be 
	using int to represent integers unless there is a need to use a specific sized integer.
size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 
	in 64 bit systems
	
-> Complex types
complex64: complex numbers which have float32 real and imaginary parts
complex128: complex numbers with float64 real and imaginary parts

The builtin function complex is used to construct a complex number with real and imaginary parts. The 
complex function has the following definition
Syntax: func complex(r, i FloatType) ComplexType  
	
*/