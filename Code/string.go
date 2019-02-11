package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    name := "Hello World"
    fmt.Printf("\n")
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)

    alphabets := "abcd xyz ABCD XYZ"
    fmt.Printf("\n\n")
    printBytes(alphabets)
    fmt.Printf("\n")
    printChars(alphabets)

    fmt.Println("\n") // The End.. .
}

func printBytes(s string) {
	fmt.Printf("HexaDecimal values..\n")  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}


func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c  ",s[i])
    }
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run string.go 
Hello Go Developer.. :]

HexaDecimal values..
48 65 6c 6c 6f 20 57 6f 72 6c 64 
H  e  l  l  o     W  o  r  l  d  

HexaDecimal values..
61 62 63 64 20 78 79 7a 20 41 42 43 44 20 58 59 5a 
a  b  c  d     x  y  z     A  B  C  D     X  Y  Z 
*/