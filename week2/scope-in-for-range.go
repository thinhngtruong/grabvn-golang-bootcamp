package week2

import (
	"fmt"
	"sync"
)

func printValue(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}
// TestScopeInForRange test
func TestScopeInForRange() {
	var wg sync.WaitGroup
	strings := []string{"a", "b", "c", "d"}
	for _, s := range strings {
		wg.Add(1)
		// bug: shared s, closure
		// go func() { 
		// 	printValue(s, &wg) 
		// }()

		// solution
		go func(s string) { 
			printValue(s, &wg) 
		}(s)
	}
	wg.Wait()
}