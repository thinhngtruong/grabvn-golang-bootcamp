package week2

import "fmt"

// PingPong example
func PingPong() {
	ping := make(chan string)
	pong := make(chan string)
	wait := make(chan bool)

	go func() {
		for v := range ping {
			fmt.Println(v)
			pong<- "pong" 
		}
	}()
	go func() {
		for v := range pong {
			fmt.Println(v)
			ping<- "ping" 
		}
	}()

	ping<- "ping"
	wait<- true
}