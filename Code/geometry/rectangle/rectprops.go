//rectprops.go

package rectangle

import (
	_ "fmt"
	"math"
)

/*
func main() {
	fmt.Println("rectangle shape properties!")
}
*/


func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}