package main


import (
	"fmt"
	)


type Income interface {
	calculate() int
	source() string
}


type FixedBilling struct {
	projectName string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours, hourlyRate int
}


func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netInome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = $%d \n", income.source(), income.calculate())
		netInome += income.calculate()
	}
	fmt.Printf("NetIncome of Organization = $%d", netInome)
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    project1 := FixedBilling{projectName: "ProjectAudi", biddedAmount: 5000}
    project2 := FixedBilling{"ProjectBMW", 8000}
    project3 := TimeAndMaterial{"ProjectCord", 130, 20}
 
 	incomeStreams := []Income { project1, project2, project3 }
 	calculateNetIncome(incomeStreams)

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/oop/PolymorphisminGo$ go run main.go
Hello Go Developer.. :]
Income from ProjectAudi = $5000 
Income from ProjectBMW = $8000 
Income from ProjectCord = $2600 
NetIncome of Organization = $15600
*/


/* => Polymorphism using interface
-> Any type which defines all the methods of an interface is said to implicitly implement that interface.
-> A variable of type interface can hold any value which implements the interface. This property of 
	interfaces is used to achieve polymorphism in Go.
*/
