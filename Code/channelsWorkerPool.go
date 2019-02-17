package main


import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	)


type Job struct {
	id, randomNo int
}
	
type Result struct {
	job Job
	sumOfDigits int
}


var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

	
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{i, randomNo}
		jobs <- job
	}
	close(jobs)
}	

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job ID: %d, input random No: %d, sum Of Digits: %d \n", 
			result.job.id, result.job.randomNo, result.sumOfDigits)
	}
	done <- true
}
	
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
	
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomNo)}
		results <- output
	}
	wg.Done()
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}


func main() {
    fmt.Println("Hello Go Developer.. :]")

	startTime := time.Now()
	
	noOfJobs := 100
	go allocate(noOfJobs)
	
	done := make(chan bool)
	go result(done)
	
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	
	<- done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken:", diff.Seconds(), "seconds.")
	
    fmt.Println("\n") // The End.. .
}


/* Output:
E:\4-Data\GoLang\Concept\Code>go run channelsWorkerPool.go
Hello Go Developer.. :]
Job ID: 2, input random No: 407, sum Of Digits: 11
Job ID: 3, input random No: 983, sum Of Digits: 20
Job ID: 4, input random No: 895, sum Of Digits: 22
Job ID: 8, input random No: 904, sum Of Digits: 13
Job ID: 0, input random No: 878, sum Of Digits: 23
Job ID: 9, input random No: 150, sum Of Digits: 6
Job ID: 6, input random No: 520, sum Of Digits: 7
Job ID: 7, input random No: 998, sum Of Digits: 26
Job ID: 5, input random No: 735, sum Of Digits: 15
Job ID: 1, input random No: 636, sum Of Digits: 15
Job ID: 12, input random No: 750, sum Of Digits: 12
Job ID: 13, input random No: 362, sum Of Digits: 11
Job ID: 11, input random No: 538, sum Of Digits: 16
Job ID: 17, input random No: 506, sum Of Digits: 11
Job ID: 16, input random No: 630, sum Of Digits: 9
Job ID: 15, input random No: 215, sum Of Digits: 8
Job ID: 14, input random No: 436, sum Of Digits: 13
Job ID: 19, input random No: 914, sum Of Digits: 14
Job ID: 10, input random No: 212, sum Of Digits: 5
Job ID: 18, input random No: 20, sum Of Digits: 2
Job ID: 23, input random No: 298, sum Of Digits: 19
Job ID: 20, input random No: 272, sum Of Digits: 11
Job ID: 25, input random No: 565, sum Of Digits: 16
Job ID: 24, input random No: 135, sum Of Digits: 9
Job ID: 21, input random No: 207, sum Of Digits: 9
Job ID: 29, input random No: 705, sum Of Digits: 12
Job ID: 26, input random No: 43, sum Of Digits: 7
Job ID: 28, input random No: 942, sum Of Digits: 15
Job ID: 22, input random No: 266, sum Of Digits: 14
Job ID: 27, input random No: 964, sum Of Digits: 19
Job ID: 38, input random No: 84, sum Of Digits: 12
Job ID: 35, input random No: 152, sum Of Digits: 8
Job ID: 31, input random No: 249, sum Of Digits: 15
Job ID: 33, input random No: 203, sum Of Digits: 5
Job ID: 34, input random No: 840, sum Of Digits: 12
Job ID: 36, input random No: 357, sum Of Digits: 15
Job ID: 37, input random No: 718, sum Of Digits: 16
Job ID: 30, input random No: 562, sum Of Digits: 13
Job ID: 32, input random No: 734, sum Of Digits: 14
Job ID: 39, input random No: 189, sum Of Digits: 18
Job ID: 48, input random No: 370, sum Of Digits: 10
Job ID: 49, input random No: 102, sum Of Digits: 3
Job ID: 47, input random No: 694, sum Of Digits: 19
Job ID: 41, input random No: 256, sum Of Digits: 13
Job ID: 46, input random No: 97, sum Of Digits: 16
Job ID: 40, input random No: 871, sum Of Digits: 16
Job ID: 42, input random No: 928, sum Of Digits: 19
Job ID: 43, input random No: 252, sum Of Digits: 9
Job ID: 44, input random No: 55, sum Of Digits: 10
Job ID: 45, input random No: 711, sum Of Digits: 9
Job ID: 58, input random No: 871, sum Of Digits: 16
Job ID: 52, input random No: 738, sum Of Digits: 18
Job ID: 54, input random No: 13, sum Of Digits: 4
Job ID: 51, input random No: 403, sum Of Digits: 7
Job ID: 57, input random No: 606, sum Of Digits: 12
Job ID: 56, input random No: 728, sum Of Digits: 17
Job ID: 55, input random No: 403, sum Of Digits: 7
Job ID: 59, input random No: 14, sum Of Digits: 5
Job ID: 50, input random No: 801, sum Of Digits: 9
Job ID: 53, input random No: 690, sum Of Digits: 15
Job ID: 60, input random No: 728, sum Of Digits: 17
Job ID: 65, input random No: 638, sum Of Digits: 17
Job ID: 69, input random No: 164, sum Of Digits: 11
Job ID: 63, input random No: 885, sum Of Digits: 21
Job ID: 62, input random No: 579, sum Of Digits: 21
Job ID: 67, input random No: 656, sum Of Digits: 17
Job ID: 66, input random No: 764, sum Of Digits: 17
Job ID: 64, input random No: 487, sum Of Digits: 19
Job ID: 61, input random No: 213, sum Of Digits: 6
Job ID: 68, input random No: 122, sum Of Digits: 5
Job ID: 73, input random No: 500, sum Of Digits: 5
Job ID: 70, input random No: 829, sum Of Digits: 19
Job ID: 79, input random No: 512, sum Of Digits: 8
Job ID: 72, input random No: 683, sum Of Digits: 17
Job ID: 75, input random No: 144, sum Of Digits: 9
Job ID: 71, input random No: 450, sum Of Digits: 9
Job ID: 77, input random No: 382, sum Of Digits: 13
Job ID: 78, input random No: 160, sum Of Digits: 7
Job ID: 74, input random No: 420, sum Of Digits: 6
Job ID: 76, input random No: 766, sum Of Digits: 19
Job ID: 83, input random No: 595, sum Of Digits: 19
Job ID: 84, input random No: 47, sum Of Digits: 11
Job ID: 81, input random No: 981, sum Of Digits: 18
Job ID: 87, input random No: 167, sum Of Digits: 14
Job ID: 80, input random No: 366, sum Of Digits: 15
Job ID: 89, input random No: 343, sum Of Digits: 10
Job ID: 88, input random No: 728, sum Of Digits: 17
Job ID: 85, input random No: 421, sum Of Digits: 7
Job ID: 86, input random No: 236, sum Of Digits: 11
Job ID: 82, input random No: 988, sum Of Digits: 25
Job ID: 91, input random No: 40, sum Of Digits: 4
Job ID: 92, input random No: 627, sum Of Digits: 15
Job ID: 98, input random No: 961, sum Of Digits: 16
Job ID: 97, input random No: 315, sum Of Digits: 9
Job ID: 96, input random No: 732, sum Of Digits: 12
Job ID: 95, input random No: 922, sum Of Digits: 13
Job ID: 94, input random No: 450, sum Of Digits: 9
Job ID: 93, input random No: 766, sum Of Digits: 19
Job ID: 90, input random No: 515, sum Of Digits: 11
Job ID: 99, input random No: 641, sum Of Digits: 11
Total time taken: 20.1069607 seconds.
*/


/* Worker Pool
-> A worker pool is a collection of threads which are waiting for tasks to be assigned to them. Once 
	they finish the task assigned, they make themselves available again for the next task.

@ref:
URL: https://golangbot.com/buffered-channels-worker-pools/
-> Topic: Worker Pool Implementation..

*/
