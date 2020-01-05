package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Work   string
	Result string
}

var wg1, wg2 sync.WaitGroup

func consume(id int, jobs <-chan *Job, results chan<- *Job) {
	defer wg1.Done()
	for job := range jobs {
		sleepMS := rand.Intn(1000)
		fmt.Printf("Worker #%d received: '%s', sleep %dms\n", id, job.Work, sleepMS)
		time.Sleep(time.Duration(sleepMS) * time.Millisecond)
		job.Result = job.Work + fmt.Sprintf("-%dms", sleepMS)
		results <- job
	}
}

func produce(jobs chan<- *Job) {
	// generating jobs
	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++
		jobs <- &Job{
			ID:   id,
			Work: fmt.Sprintf("%c", c),
		}
	}
	close(jobs)
}

func analyze(results <-chan *Job) {
	defer wg2.Done()
	for job := range results {
		fmt.Printf("result %s\n", job.Result)
	}
}

func main() {
	jobs := make(chan *Job, 100)
	results := make(chan *Job, 100)

	// start consumers
	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go consume(i, jobs, results)
	}

	// start producing
	go produce(jobs)

	// analyze results
	wg2.Add(1)
	go analyze(results)

	wg1.Wait()
	close(results)
	wg2.Wait()
}
