package gojel

import (
	"fmt"
	"time"
)

func SetTimeout(callback func(), millis int64) (chan bool, *time.Timer) {
	timer := time.NewTimer(time.Duration(millis * int64(time.Millisecond)))
	done := make(chan bool)
	go func() {
		<-timer.C
		callback()
		done <- true
	}()
	return done, timer
}

func SetInterval(callback func(), millis int64) (chan bool, *time.Ticker) {
	ticker := time.NewTicker(time.Duration(millis * int64(time.Millisecond)))
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				callback()
			}
		}
	}()
	return done, ticker
}

func jobWorker(callback func(interface{}) interface{}, counter *int, id int, jobs <-chan interface{}, results chan<- JobResult) {
	for j := range jobs {
		*counter++
		var jobNumber = *counter
		fmt.Printf("Worker %v: started job %v with value: %v\n", id, jobNumber, j)
		result := callback(j)
		results <- JobResult{
			WorkerNum: id,
			Result:    result,
		}
		fmt.Printf("Worker %v: finished job %v\n", id, jobNumber)
	}
}

type JobResult struct {
	WorkerNum int
	Result    interface{}
}

func StartJobs(callback func(interface{}) interface{}, workersNum int) (chan<- interface{}, <-chan JobResult) {
	jobs := make(chan interface{}, workersNum)
	results := make(chan JobResult, workersNum)
	counter := 0
	for w := 1; w <= workersNum; w++ {
		go jobWorker(callback, &counter, w, jobs, results)
	}
	return jobs, results
}
