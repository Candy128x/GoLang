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

	a := 23
	fmt.Println("type of a is: %T",a)
	fmt.Println("size of a is: %d", unsafe.Sizeof(a))
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
*/