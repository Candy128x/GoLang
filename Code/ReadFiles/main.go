package main


import (
	"fmt"
	"io/ioutil"
	"flag"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    /* -> Reading an entire file into memory */
    // data, err := ioutil.ReadFile("note.txt")
    // if err != nil {
    // 	fmt.Println("File reading error !")
    // 	return
    // }
    // fmt.Println("Content of file =", string(data))

    /* Output: 
	[ERROR] if set ReadFile("note3.txt")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
	Hello Go Developer.. :]
	File reading error !

    [CORRECT] if set ReadFile("note.txt")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
	Hello Go Developer.. :]
	Content of file = Go has concurrency feature &
	also channel feature :)
	*/


	/* -> 1. Using absolute file path */
    // data, err := ioutil.ReadFile("/home/ashishs/Downloads/storage/Go/GoLang/Code/ReadFiles/note.txt")
    // if err != nil {
    // 	fmt.Println("File reading error !")
    // 	return
    // }
    // fmt.Println("Content of file =", string(data))

    /* Output:
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
	Hello Go Developer.. :]
	Content of file = Go has concurrency feature &
	also channel feature :)
	*/


	/* -> 2. Passing the file path as a command line flag */
	fptr := flag.String("fpath", "note.txt", "file path to read from")
	flag.Parse()
	fmt.Println("Value of fpath is =", *fptr)
	data, err := ioutil.ReadFile(*fptr)
    if err != nil {
    	fmt.Println("File reading error!")
    	return
    }
    fmt.Println("Content of file =", string(data))


	/* Output:
	[CORRECT] if set -> flag.String("fpath", "note.txt", "file path to read from")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
	Hello Go Developer.. :]
	Value of fpath is = note.txt
	Content of file = Go has concurrency feature &
	also channel feature :)

	[ERROR] if set -> flag.String("fpath", "note3.txt", "file path to read from")
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
	Hello Go Developer.. :]
	Value of fpath is = note3.txt
	File reading error!
	*/


	/* -> 3. Bundling the text file along with the binary */


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
Hello Go Developer.. :]
Value of fpath is = note.text
*/


/* Title
-> Description..
- Reading an entire file into memory
	- Using absolute file path
	- Passing the file path as a command line flag
	- Bundling the file inside the binary
- Reading a file in small chunks
- Reading a file line by line

=> The reason is Go is a compiled language. What go install does is, it creates a binary from the 
	source code. The binary is independent of the source code and it can be run from any 
	location. Since test.txt is not found in the location from which the binary is run, the 
	program complains that it cannot find the file specified.
-> There are three ways to approach this problem,
	1. Using absolute file path
	2. Passing the file path as a command line flag
	3. Bundling the text file along with the binary
*/
