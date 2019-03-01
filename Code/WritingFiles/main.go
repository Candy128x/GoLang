package main


import (
	"fmt"
	"os"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")


    // f, err := os.Create("note.txt") // -> This function returns a File descriptor.
    // if err != nil {
    // 	fmt.Println("File creation err:", err)
    // 	return
    // }

    // l, err := f.WriteString("Hey dude ~") // -> This method returns the number of bytes written and error if any.
    // if err != nil {
    // 	fmt.Println("File writing err:", err)
    // 	f.Close()
    // 	return
    // }

    // fmt.Println(l, ", bytes written successfully :)")
    
    // err = f.Close()
    // if err != nil {
    // 	fmt.Println("File closing err,", err)
    // 	return
    // }

    /* Output:
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/WritingFiles$ go run main.go 
	Hello Go Developer.. :]
	10 , bytes written successfully :)

	=> note.txt file
	-> Hey dude ~
	*/


    /* -> Writing bytes to a file */
    // f, err := os.Create("note.txt")
    // if err != nil {
    // 	fmt.Println("File creation err:", err)
    // 	return
    // }

    // d2 := []byte{23, 34, 101, 82, 69, 11, 07, 54, 73}

    // n2, err := f.Write(d2)
    // if err != nil {
    // 	fmt.Println("File writing err:", err)
    // 	f.Close()
    // 	return
    // }

    // fmt.Println(n2, ", bytes written successfully :)")
    
    // err = f.Close()
    // if err != nil {
    // 	fmt.Println("File closing err,", err)
    // 	return
    // }
    
    /* Output:
	ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/WritingFiles$ go run main.go 
	Hello Go Developer.. :]
	9 , bytes written successfully :)

	=> note.txt file
	-> "eRE6I
	*/


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/WritingFiles$ go run main.go 
Hello Go Developer.. :]
9 , bytes written successfully :)
*/


/* Title
-> Description..

=> Operation
- Writing string to a file
- Writing bytes to a file
- Writing data line by line to a file
- Appending to a file
- Writing to file concurrently
*/
