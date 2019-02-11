package main


import (
	"fmt"
	)


func main() {
    fmt.Println("Hello Go Developer.. :]")

    var arr1 [5]int //int array with length 5
    fmt.Println(arr1)	

    var arr2 [5]string
    fmt.Println(arr2)	

    arr1[0] = 12
    arr1[1] = 13
    arr1[2] = 23
    arr1[3] = 18
    arr1[4] = 29
	fmt.Println(arr1)

	arr3 := [5]int{11, 34, 23, 22, 43}
	fmt.Println(arr3)

	arr4 := [...]int {21, 32, 31, 21, 32, 54, 44}
	fmt.Println(arr4)
	fmt.Println("\narr4:", arr4)
	fmt.Println("arr4 len:", len(arr4))
	fmt.Println("arr4 cap:", cap(arr4))
	
	arr4Slice := arr4[2:5]
	fmt.Println("\narr4Slice:", arr4Slice)
	fmt.Println("arr4Slice len:", len(arr4Slice))
	fmt.Println("arr4Slice cap:", cap(arr4Slice))

	a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b)

    fmt.Println("arr3 is before replace zero th position", arr3)
    arrReplaceZerothPosition(arr3)
    fmt.Println("arr3 is after passing to function..", arr3)
    /* Output:
	arr3 is before replace zero th position [11 34 23 22 43]
	inside func:  [3 34 23 22 43]
	arr3 is after passing to function.. [11 34 23 22 43]
    */

	// slice
	fmt.Println("arr3 =>", arr3)
	fmt.Println("arr3[2:] =>", arr3[2:])
	fmt.Println("arr3[:3] =>", arr3[:3])
	fmt.Println("arr3[1:3] =>", arr3[1:3])

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water-melon", "pine-apple", "chikoo"}
	fmt.Println("\n fruitarray =>", fruitarray)
    fruitslice := fruitarray[1:3]
    fmt.Println("\n fruitslice =>", fruitslice)
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))

    //append
    cars := []string{"Ferrari", "Honda", "Ford"}
    cars = append(cars, "Toyota")
	fmt.Println("\n fruitarray =>", cars)


	var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("\n slice is nil going to append")
        names = append(names, "john", "shawn", "smith")
        fmt.Println("names contents:",names)
    }

    
    veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
    fmt.Println("\n food:",food)


    fmt.Println("\n\n") // The End.. .
}


func arrReplaceZerothPosition(arr3 [5]int){
	arr3[0] = 3
	fmt.Println("inside func: ", arr3)
}

/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run array.go 
Hello Go Developer.. :]
[0 0 0 0 0]
[    ]
[12 13 23 18 29]
[11 34 23 22 43]
[21 32 31 21 32 54 44]

arr4: [21 32 31 21 32 54 44]
arr4 len: 7
arr4 cap: 7

arr4Slice: [31 21 32]
arr4Slice len: 3
arr4Slice cap: 5
a is  [USA China India Germany France]
b is  [Singapore China India Germany France]
arr3 is before replace zero th position [11 34 23 22 43]
inside func:  [3 34 23 22 43]
arr3 is after passing to function.. [11 34 23 22 43]
arr3 => [11 34 23 22 43]
arr3[2:] => [23 22 43]
arr3[:3] => [11 34 23]
arr3[1:3] => [34 23]

 fruitarray => [apple orange grape mango water-melon pine-apple chikoo]

 fruitslice => [orange grape]
length of slice 2 capacity 6
After re-slicing length is 6 and capacity is 6

 fruitarray => [Ferrari Honda Ford Toyota]

 slice is nil going to append
names contents: [john shawn smith]

 food: [potatoes tomatoes brinjal oranges apples]
*/