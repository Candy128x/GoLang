//geometry.go

package main

import (
	"fmt"
	"geometry/rectangle" //import custom package
)

func main() {  
    var rectLen, rectWidth float64 = 6, 7
    fmt.Println("Geometrical shape properties")
        /*Area function of rectangle package used
        */
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
        /*Diagonal function of rectangle package used
        */
    fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}

/* Output:
E:\4-Data\GoLang\Concept\Code\geometry>go run geometry.go
Geometrical shape properties
area of rectangle 42.00
diagonal of the rectangle 9.22
*/