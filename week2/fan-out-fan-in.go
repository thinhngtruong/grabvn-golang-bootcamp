package week2

import (
	"fmt"
	"sync"
)

var x = make(chan int)
var x1 = make(chan int)
var x2 = make(chan int)
var x3 = make(chan int)
var out = make(chan int)

var total int
var waitgroup sync.WaitGroup

func spread(x <-chan int, x1 chan<- int, x2 chan<- int, x3 chan<- int) {
	for value := range x {
		x1 <- value
		x2 <- value
		x3 <- value
	}
}

func passCalculatedValue(ch <-chan int, f func(int) int, out chan<- int) {
	for value := range ch {
		out<- f(value)
	}
}

func sum(ch <-chan int, wg *sync.WaitGroup) {
	for value := range ch {
		total += value
		wg.Done()
	}
}

// FanoutFanin pass integer numbers from 1 to 10 to channel x.
// Read values from channel x and pass recieved value to 3 channels x1, x2, x3
// Read values from 3 channels x1, x2, x3 and pass calculated value to channel out
// Read all values of channel out and calculate sum
func FanoutFanin() {
	go spread(x, x1, x2, x3)
	go passCalculatedValue(x1, func(value int) int { return value }, out)
	go passCalculatedValue(x2, func(value int) int { return value*2 }, out)
	go passCalculatedValue(x3, func(value int) int { return value*3 }, out)
	go sum(out, &waitgroup)

	for i := 1; i <= 10; i++ {
		waitgroup.Add(3)
		x <- i
	}
	waitgroup.Wait()

	fmt.Println("Total = ", total)
}