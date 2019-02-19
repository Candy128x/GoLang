package main


import (
	"fmt"
	"os"
	_"net"
	_"path/filepath"
)


func main() {
	fmt.Println("Hello Go Developer.. :]")

	f, err := os.Open("/note.txt") //return the file handler and error 
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "-> Open Successfully :)")

	fmt.Printf("\n vari f \t| value = %s \t| type = %T", f, f)
	fmt.Printf("\n vari err \t| value = %s \t\t\t\t\t| type = %T", err, err)

	/* Output:
	[err] if wrote -> os.Open("/note.txt") -> file not exist!
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
	Hello Go Developer.. :]
	open /note.txt: no such file or directory


	[correct] if wrote -> os.Open("note.txt")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
	Hello Go Developer.. :]
	note.txt -> Open Successfully :)

	vari f 	| value = &{%!s(*os.file=&{3 note.txt <nil> 0})} 	| type = *os.File
	vari err 	| value = %!s(<nil>) 								| type = <nil>
	*/


	// => Different ways to extract more information from the error

	/* -> 1. Asserting the underlying struct type and getting more information from the struct fields */
	// f2, err2 := os.Open("/note3.txt")
	// if err2, ok := err2.(*os.PathError); ok {
	// 	fmt.Println("File2 at path=", err2.Path, "failed to open :(")
	// 	return
	// }
	// fmt.Println(f2.Name(), "-> Open2 Successfully :)")

	/* Output:
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
	Hello Go Developer.. :]
	File2 at path= /note3.txt failed to open :(	
	*/


	/* -> 2. Asserting the underlying struct type and getting more information using methods */
	// addr3, err3 := net.LookupHost("golang12345.com")
	// if err3, ok3 := err3.(*net.DNSError); ok3 {
	// 	if err3.Timeout() {
	// 		fmt.Println("Operation time out.. .!")
	// 	} else if err3.Temporary() {
	// 		fmt.Println("Temporary error.. !")
	// 	} else {
	// 		fmt.Println("Generic error! =", err3)
	// 	}
	// 	return
	// }
	// fmt.Println("Address=",addr3)

	/* Output:
	[correct] if net.LookupHost("golang123.com")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
	Hello Go Developer.. :]
	Address= [42.121.236.221]

	[err] if net.LookupHost("golang12345.com")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
	Hello Go Developer.. :]
	Generic error! = lookup golang12345.com: no such host
	*/


	/* -> 3. Direct comparison */
	// files4, err4 := filepath.Glob("[") // The Glob function of the filepath package is used to return the names of all files that matches a pattern. This function returns an error ErrBadPattern when the pattern is malformed.
	// if err4 != nil && err4 == filepath.ErrBadPattern {
	// 	fmt.Println(err4)
	// 	return
	// }
	// fmt.Println("Matched files =",files4)

	/* Output:
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
	Hello Go Developer.. :]
	syntax error in pattern
	*/


	fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/errorHandling$ go run error.go 
Hello Go Developer.. :]
syntax error in pattern
*/


/* => What are errors?
-> Errors indicate an abnormal condition in the program. Let's say we are trying to open a file and the file 
does not exist in the file system. This is an abnormal condition and it's represented as an error.

=> func Open(name string) (file *File, err error)
-> If the file has been opened successfully, then the Open function will return the file handler and 
error will be nil. If there is an error while opening the file, a non nil error will be returned.
-> The idiomatic way of handling error in Go is to compare the returned error to nil. A nil value 
indicates that no error has occurred and a non nil value indicates the presence of an error.
*/
