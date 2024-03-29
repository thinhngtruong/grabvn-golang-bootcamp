package week2

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
    out := make(chan int, len(nums))
    for _, n := range nums {
        out <- n
    }
    close(out)
    return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
    defer close(out)
    out := make(chan int)
    go func() {
        for n := range in {
            select {
            case out <- n * n:
            case <-done:
                return
            }
        }
        close(out)
    }()
    return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c or done is closed, then calls
    // wg.Done.
    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            select {
            case out <- n:
            case <-done:
                // when done is closed, <-done immidiately return zero value of channel type
                return
            }
        }
    }
	
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

// FanoutFaninWithPipeline test
func FanoutFaninWithPipeline() {
    // Set up a done channel that's shared by the whole pipeline,
    // and close that channel when this pipeline exits, as a signal
    // for all the goroutines we started to exit.
    done := make(chan struct{})
    defer close(done)
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	// so values of in either read from c1 or c2
    c1 := sq(done, in)
    c2 := sq(done, in)

    // Consume the first value from output.
    out := merge(done, c1, c2)
    fmt.Println(<-out) // 4 or 9

    // done will be closed by the deferred call.
}