package main

import (
	"fmt"
	"math/rand"
	"time"

	gojel "github.com/sandrolain/go-jel"
)

func main() {
	// Start the pool of workers
	jobs, results := gojel.StartJobs(func(value interface{}) interface{} {
		fmt.Printf("Worker Input %v\n", value)
		time.Sleep(time.Millisecond * 100)
		return value.(int) * 2
	}, 3)

	go func() {
		for res := range results {
			fmt.Printf("Worker Result %v\n", res)
		}
	}()

	counter := 0
	itvStop, _ := gojel.SetInterval(func() {
		counter++
		fmt.Printf("\nInterval %v\n", counter)
		jobs <- rand.Intn(100)
	}, 250)
	gojel.SetTimeout(func() {
		itvStop <- true
		fmt.Printf("Stop Interval!\n")
	}, 3000)
	toDone, _ := gojel.SetTimeout(func() {
		fmt.Printf("End!\n")
	}, 5000)
	<-toDone
}
