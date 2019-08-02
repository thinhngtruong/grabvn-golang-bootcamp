// Package workerpool inspire by https://github.com/dmora/workerpool/blob/master/pool.go
package workerpool

import (
	"sync"
	"time"
	"fmt"
)

// ProcessorFunc signature that defines the dependency injection to process "Jobs"
type ProcessorFunc func(resource interface{}) error

// ResultProcessorFunc signature that defines the dependency injection to process "Results"
type ResultProcessorFunc func(result Result) error

// Job Structure that wraps Jobs information
type Job struct {
	id       int
	resource interface{}
}

// Result holds the main structure for worker processed job results.
type Result struct {
	Job Job
	Err error
}

// Pool generic struct that keeps all the logic to manage the queues
type Pool struct {
	numRoutines int
	jobs        chan Job
	results     chan Result
	done        chan bool
	completed   bool
}

// NewPool returns a new manager structure ready to be used.
func NewPool(numRoutines int) *Pool {
	p := &Pool{numRoutines: numRoutines}
	p.jobs = make(chan Job, numRoutines)
	p.results = make(chan Result, numRoutines)
	p.done = make(chan bool)
	// fmt.Println("[NewPool] Created a new Pool")
	return p
}

// Start start a worker pool
func (p *Pool) Start(resources []interface{}, procFunc ProcessorFunc, resFunc ResultProcessorFunc) {
	fmt.Println("[Start] worker pool starting")
	startTime := time.Now()
	go p.allocate(resources)
	go p.collect(resFunc)
	go p.workerPool(procFunc)
	<-p.done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("[Start] worker pool end, total time taken: [%f] seconds\n", diff.Seconds())
}

// allocate allocates jobs based on an array of resources to be processed by the worker pool
func (p *Pool) allocate(resources []interface{}) {
	defer close(p.jobs)
	// fmt.Printf("[allocate] Start allocating [%d] resources\n", len(resources))
	for i, v := range resources {
		job := Job{id: i+1, resource: v}
		p.jobs <- job
	}
}

// work performs the actual work by calling the processor and passing in the Job as reference obtained
// from iterating over the "Jobs" channel
func (p *Pool) work(wg *sync.WaitGroup, processor ProcessorFunc) {
	defer wg.Done()
	for job := range p.jobs {
		// fmt.Printf("[work] working on Job ID [%d]\n", job.id)
		output := Result{job, processor(job.resource)}
		// fmt.Printf("[work] done with Job ID [%d]\n", job.id)
		p.results <- output
	}
}

// workerPool creates or spawns new "work" goRoutines to process the "Jobs" channel
func (p *Pool) workerPool(processor ProcessorFunc) {
	defer close(p.results)
	var wg sync.WaitGroup
	for i := 0; i < p.numRoutines; i++ {
		wg.Add(1)
		go p.work(&wg, processor)
		// fmt.Printf("[workerPool] Spawned work goRoutine [%d]\n", i+1)
	}
	wg.Wait()
	// fmt.Println("[workerPool] all work goroutines done processing")
}

// collect post processes the channel "Results" and calls the ResultProcessorFunc passed in as reference further processing.
func (p *Pool) collect(proc ResultProcessorFunc) {
	for result := range p.results {
		outcome := proc(result)
		fmt.Printf("[collect] Job with id: [%d] completed, outcome: %s\n", result.Job.id, outcome)
	}
	p.done <- true
	p.completed = true
}

// IsCompleted utility method to check if all work has done from an outside caller.
func (p *Pool) IsCompleted() bool {
	return p.completed
}
