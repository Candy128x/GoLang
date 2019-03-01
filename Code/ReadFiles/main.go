package main


import (
	"fmt"
	_"io/ioutil"
	"flag"
	"os"
	"log"
	"bufio"
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
	// fptr := flag.String("fpath", "note.txt", "file path to read from") // ->  This function returns the address of the string variable that stores the value of the flag.
	// flag.Parse()
	// fmt.Println("Value of fpath is =", *fptr)
	// data, err := ioutil.ReadFile(*fptr)
 //    if err != nil {
 //    	fmt.Println("File reading error!")
 //    	return
 //    }
 //    fmt.Println("Content of file =", string(data))


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


	/* -> Reading file in small chunks*/
	// fptr := flag.String("fpath", "note.txt", "file path to read from")
	// flag.Parse()

	// f, err := os.Open(*fptr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	if err = f.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// r := bufio.NewReader(f)
	// b := make([]byte, 3)

	// for {
	// 	_, err := r.Read(b)
	// 	if err != nil {
	// 		fmt.Println("Error Reading File :", err)
	// 		break
	// 	}
	// 	fmt.Println(string(b))
	// }

	/* Output:
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
	Hello Go Developer.. :]
	Go 
	has
	 co
	ncu
	rre
	ncy
	 fe
	atu
	re 
	&
	a
	lso
	 ch
	ann
	el 
	fea
	tur
	e :
	)
	:
	Error Reading File : EOF
	*/


	/* -> Rading a file line by line */
	fptr := flag.String("fpath", "note.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/ReadFiles$ go run main.go 
Hello Go Developer.. :]
Go 
has
 co
ncu
rre
ncy
 fe
atu
re 
&
a
lso
 ch
ann
el 
fea
tur
e :
)
:
Error Reading File : EOF
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

=> how the flag package works. 
-> The flag package has a String function. This function accepts 3 arguments. The first is the name 
	of the flag, second is the default value and the third is a short description of the flag.	
-> flag.Parse() should be called before any flag is accessed by the program.
*/
