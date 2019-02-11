package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    var Salary map[string]int
    if Salary == nil {
    	fmt.Println("Salary is nil")
    }


    mapFruits := map[string]int {
    	"grapes" : 60,
    	"pinapple" : 80,
    }
    fmt.Println("map mapFruits values: ", mapFruits)

    mapFruits["apple"] = 32
    mapFruits["banana"] = 20
    mapFruits["chikoo"] = 10
    fmt.Println("map mapFruits values: ", mapFruits)

    fmt.Println("map mapFruits values of apple :", mapFruits["apple"])

    // newFruit := "cherry"
    newFruit := "banana"
    value, ok := mapFruits[newFruit]
    if ok == true{
    	fmt.Println("Price of,", newFruit, "is Rs.", value)
    } else {
    	fmt.Println(newFruit, "not found!")
    }


    fmt.Println("All Fruit Prices..")
    for key, value := range mapFruits {
        fmt.Printf("mapFruits[%s] = %d\n", key, value)
    }

    // Delete item
    fmt.Println("--deleting item--")
    fmt.Println("map mapFruits before delete values: ", mapFruits)
    delete(mapFruits, "pinapple")
    fmt.Println("map mapFruits after delete values: ", mapFruits)


    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run maps.go 
Hello Go Developer.. :]
Salary is nil
map mapFruits values:  map[grapes:60 pinapple:80]
map mapFruits values:  map[grapes:60 pinapple:80 apple:32 banana:20 chikoo:10]
map mapFruits values of apple : 32
Price of, banana is Rs. 20
All Fruit Prices..
mapFruits[grapes] = 60
mapFruits[pinapple] = 80
mapFruits[apple] = 32
mapFruits[banana] = 20
mapFruits[chikoo] = 10
--deleting item--
map mapFruits before delete values:  map[grapes:60 pinapple:80 apple:32 banana:20 chikoo:10]
map mapFruits after delete values:  map[grapes:60 apple:32 banana:20 chikoo:10]
*/


/*
-> What is a map?
- A map is a builtin type in Go which associates a value to a key. The 
	value can be retrieved using the corresponding key.
*/