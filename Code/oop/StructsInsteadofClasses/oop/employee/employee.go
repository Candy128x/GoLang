package employee


import (
	"fmt"
	)


type employee struct {
	firstName, lastName string
	totalLeaves, leavesTaken int
}

// New() function instead of constructors
func New(firstName, lastName string, totalLeaves, leavesTaken int) employee {
	e := employee {firstName, lastName, totalLeaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining.", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run Demo.go 
Hello Go Developer.. :]
*/


/* Title
-> Description..

@note
=> ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code/oop/StructsInsteadofClasses$ sudo cp -r oop /usr/lib/go/src/pkg/
*/
