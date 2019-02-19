package main

import (
    "fmt"
    "os"
    )

func main() {
    fmt.Println("GoPath: ", os.Getenv("GOPATH"))
}