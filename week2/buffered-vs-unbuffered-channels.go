package week2

import "fmt"

// BufferedVsUnbufferedChannels size 1 vs size 0
func BufferedVsUnbufferedChannels() {
	// TODO: confirm this

	// buffered channel, buffered size = 1
	ch := make(chan bool, 1)
	ch<- true // not lock the current gorutine
	fmt.Println("hello")
	fmt.Println(<-ch)

	// unbuffered channel, buffered size = 0
	// ch := make(chan bool)
	// ch<- true // lock the current gorutine, DEADLOCK!
	// fmt.Println("hello")
	// fmt.Println(<-ch)
}