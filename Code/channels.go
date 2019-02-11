package main


import (
	"fmt"
	"time"
	)


func hello(done chan bool) {
	fmt.Println("Hello World goRoutine")
	done <- true
}


func helloSleep(doneSleep chan bool) {
	fmt.Println("In side helloSleep, helloSleep is going to sleep..")
	time.Sleep(5 * time.Second)
	fmt.Println("helloSleep go routines awake & going to write to done..")
	doneSleep <- true
}


func sendData(sendch chan<- int) { // unidirectional | received only
	sendch <- 10 // write
	fmt.Println("in sendData goRoutnies..")
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

    var a chan int // it's nil
    fmt.Println("Values of channel a:", a)
    // op: Values of channel a: <nil>


    /*
    a := make(chan int)
    // Sending & Receiving form channel
    a <- data // write to a channel a
    data := <- a // read from a channel a
	*/


    done := make(chan bool)
    go hello(done)
    <- done
    fmt.Println("--main function--")

    /* op:
    ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run channels.go 
	Hello Go Developer.. :]
	Values of channel a: <nil>
	Hello World goRoutine
	--main function--
	*/


	// => Deadlock..
	// ch := make(chan int)
	// ch <- 5
	/* explain:
	- In the program above, a channel ch is created and we send 5 to the channel in line ch <- 5. 
	- In this program no other Goroutine is receiving data from the channel ch. 
	- Hence this program will panic with the following runtime error.
	
	=> err:
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.main()
		/home/ashishs/Downloads/storage/Go/GoLang/Code/channels.go:46 +0x2d0
	exit status 2
	*/


	// sleep/awake features
	doneSleep := make(chan bool)
	fmt.Println("--Main going to call HelloSleep  go goRoutine..")
	go helloSleep(doneSleep)
	<- doneSleep
	fmt.Println("--Main received data--")
	/* op:
	Hello World goRoutine
	--main function--
	--Main going to call HelloSleep  go goRoutine..
	In side helloSleep, helloSleep is going to sleep..
	helloSleep go routines awake & going to write to done..
	--Main received data--

	=> Explanation
	- In the above program we have introduced a sleep of 5 seconds to the hello function in 
		line no. 10.
	- This program will first print Main going to call hello go goroutine. Then the hello Goroutine 
		will be started and it will print hello go routine is going to sleep. After this is printed, 
		the hello Goroutine will sleep for 5 seconds and during this time main Goroutine will be 
		blocked since it is waiting for data from the done channel in line <-done. After 5 seconds 
		hello go routine awake and going to write to done will be printed followed by Main received data.
	*/


	// => Unidirectional channels
	fmt.Println("--Unidirectional channels <- send only channels.. ")
	/* All the channels we discussed so far are bidirectional channels, that is data can be both 
		sent and received on them. It is also possible to create unidirectional channels, that is 
		channels that only send or receive data.
	*/
	// chan<- int denotes a send only channel
	// sendch := make(chan<- int) // unidirectional
	sendch := make(chan int) // bidirectional
	go sendData(sendch)
	fmt.Println("Print/Read <-sendch: ",<-sendch)

	// fmt.Println(<-sendch) // read
	/* err:
	# command-line-arguments
	./channels.go:110: invalid operation: <-sendch (receive from send-only type chan<- int)
	*/

    fmt.Println("\n") // The End.. .
}


/* Output:
ashishs@lp-0731:~/Downloads/storage/Go/GoLang/Code$ go run pointers.go 
Hello Go Developer.. :]

*/

/* What are channels
-> Channels can be thought as pipes using which Goroutines communicate. Similar to how water flows from 
	one end to another in a pipe, data can be sent from one end and received from the another end 
	using channels.

=> Sends and receives are blocking by default
-> Sends and receives to a channel are blocking by default. What does this mean? When a data is sent to a 
	channel, the control is blocked in the send statement until some other Goroutine reads from that 
	channel. Similarly when data is read from a channel, the read is blocked until some Goroutine writes 
	data to that channel.

=> Deadlock
-> One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a 
	channel, then it is expected that some other Goroutine should be receiving the data. If this does not 
	happen, then the program will panic at runtime with Deadlock.
*/