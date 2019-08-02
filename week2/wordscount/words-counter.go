package wordscount

import (
	"fmt"
	"github.com/nhaancs/grabvn-golang-bootcamp/week2/workerpool"
	"io/ioutil"
	"strings"
	"sync"
	"net/http"
)

// WordsCount count occurrences of words from files, urls, directory's files or list of 3 kinds of them.
// User input through command line. Input values are separated by space.
// Only process 10 resources at a time (use worker pool).
func WordsCount() {
	var mutex = &sync.Mutex{}
	var counter = make(map[string]int)
	var inputs = readInputParams()
	var paths = getValidPaths(inputs)

	resources := make([]interface{}, len(paths))
	for i, path := range paths {
		resources[i] = path
	}
	pool := workerpool.NewPool(numWorkers)
	pool.Start(resources, count(counter, mutex), resultsCollector)

	fmt.Println(counter)
}

func count(counter map[string]int, mutex *sync.Mutex) workerpool.ProcessorFunc {
	resourceProcessor := func (resource interface{}) error {
		path := resource.(string)

		if isValidURL(path) {
			return countWithURL(counter, path, mutex)
		}

		return countWithFile(counter, path, mutex)
	}

	return resourceProcessor
}

func countWithFile(counter map[string]int, path string, mutex *sync.Mutex) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	countWord(counter, content, mutex)
	return nil
}

func countWithURL(counter map[string]int, url string, mutex *sync.Mutex) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	countWord(counter, html, mutex)
	return nil
}

func countWord(counter map[string]int, content []byte, mutex *sync.Mutex) {
	contentSlice := strings.Fields(strings.ToLower(string(content)))
	for _, word := range contentSlice {
		mutex.Lock()
		counter[word]++
		mutex.Unlock()
	}
}

func resultsCollector(result workerpool.Result) error {
	if result.Err != nil {
		fmt.Printf("[resultProcessor] Got error: %s\n", result.Err)
	}
	return nil
}
