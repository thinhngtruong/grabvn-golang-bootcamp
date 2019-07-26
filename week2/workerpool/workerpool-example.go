package workerpool

import (
	"fmt"
)

func resourceProcessor(resource interface{}) error {
	fmt.Printf("[resourceProcessor] Resource processor got: %s\n", resource)
	return nil
}

func resultProcessor(result Result) error {
	fmt.Printf("[resultProcessor] Result processor got: %s\n", result.Err)
	return nil
}

// ExamplePoolStart Example
func ExamplePoolStart() {
	strings := []string{"first", "second"}
	resources := make([]interface{}, len(strings))
	for i, s := range strings {
		resources[i] = s
	}

	pool := NewPool(3)
	pool.Start(resources, resourceProcessor, resultProcessor)
}