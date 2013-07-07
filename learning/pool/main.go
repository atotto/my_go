package main

import (
	"fmt"
	"runtime"
	//"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)

		// do something
		//time.Sleep(time.Second)
		for i := 0; i < 10000000; i++ {
		}

		// send result
		results <- j * j
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 1; j <= 100; j++ {
		<-results
	}
}
