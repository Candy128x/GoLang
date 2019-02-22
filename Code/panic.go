package main


import (
	"fmt"
	)


func fullName(firstName string, lastName string) {
	if firstName == "" {
		panic("Run Time err!, FirstName cann't be blank :(")
	}
	if lastName == "" {
		panic("Run Time err!, LastName cann't be blank :(")
	}
	fmt.Printf("FullName is = %s %s \n", firstName, lastName)
	fmt.Println("Returned normally from fullName() :)")
}



func main() {
    fmt.Println("Hello Go Developer.. :]")

    firstName := "Shamu"
    fullName(firstName, "")
    fmt.Println("Returned normally from main() :)")


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run panic.go 
Hello Go Developer.. :]
panic: Run Time err!, LastName cann't be blank :(

goroutine 1 [running]:
runtime.panic(0x47fe00, 0xc21000a170)
	/usr/lib/go/src/pkg/runtime/panic.c:266 +0xb6
main.fullName(0xc21000a160, 0x0)
	/home/ashishs/Downloads/storage/Go/GoLang/Code/panic.go:14 +0xbb
main.main()
	/home/ashishs/Downloads/storage/Go/GoLang/Code/panic.go:26 +0xe4
exit status 2
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ 

@note
- On run time error it's print panic statment as well system/compiler error message.																																									 
*/


/* => Panic!
-> We use panic to terminate the execution of the program.
-> When a function encounters a panic, its execution is stopped, any deferred functions are 
	executed and then the control returns to its caller.
-> This process continues until all the functions of the current goroutine have returned at 
	which point the program prints the panic message, followed by the stack trace and then 
	terminates. This concept will be more clear when we write a sample program.
-> panic and recover can be considered similar to try-catch-finally idiom in other languages 
	except that it is rarely used and when used is more elegant and results in clean code.

=> When should panic be used?
-> One important factor is that you should avoid panic and recover and use errors where ever 
	possible. Only in cases where the program just cannot continue execution should a panic and 
	recover mechanism be used.

=> There are two valid use cases for panic.
1) An unrecoverable error where the program cannot simply continue its execution. 
	One example would be a web server which fails to bind to the required port. In this case 
	it's reasonable to panic as there is nothing else to do if the port binding itself fails.
2) A programmer error. 
	Let's say we have a method which accepts a pointer as a parameter and someone calls this 
	method using nil as argument. In this case we can panic as it's a programmer error to 
	call a method with nil argument which was expecting a valid pointer.
*/
