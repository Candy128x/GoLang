package main


import (
	"fmt"
	"runtime/debug"
	)

	
func a() {
	defer r()
	n := []int {5, 7, 4}
	fmt.Println("Value of array n =", n[3])
	fmt.Println("Normally returned from a() :)")
}

func r() {
	if r := recover(); r != nil {
	fmt.Println("Recovered =", r)
	debug.PrintStack()
	}
}
	

func main() {
    fmt.Println("Hello Go Developer.. :]")
	
	a()
	fmt.Println("Normally returned from main() :)")
	
    fmt.Println("\n") // The End.. .
}


/* Output:
--1--
-> without import "runtime/debug" & func r() -> debug.PrintStack()
E:\4-Data\GoLang\Concept\Code>go run panicStackTrace.go
Hello Go Developer.. :]
Recovered = runtime error: index out of range
Normally returned from main() :)


--2--
E:\4-Data\GoLang\Concept\Code>go run panicStackTrace.go
Hello Go Developer.. :]
Recovered = runtime error: index out of range
goroutine 1 [running]:
runtime/debug.Stack(0xc000074008, 0xc000071de8, 0x2)
        C:/Go/src/runtime/debug/stack.go:24 +0xae
runtime/debug.PrintStack()
        C:/Go/src/runtime/debug/stack.go:16 +0x29
main.r()
        E:/4-Data/GoLang/Concept/Code/panicStackTrace.go:20 +0xa3
panic(0x4a6da0, 0x555400)
        C:/Go/src/runtime/panic.go:513 +0x1c7
main.a()
        E:/4-Data/GoLang/Concept/Code/panicStackTrace.go:13 +0x45
main.main()
        E:/4-Data/GoLang/Concept/Code/panicStackTrace.go:28 +0x6d
Normally returned from main() :)
*/


/* => Getting stack trace after recover
-> If we recover a panic, we lose the stack trace about the panic. Even in the program 
	above after recovery, we lost the stack trace.
-> There is a way to print the stack trace using the PrintStack function of the Debug package.
*/
