package wordscount

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func getValidPaths(inputs []string) []string {
	var paths []string
	for _, input := range inputs {
		if isValidFileOrDirectoryPath(input) {
			if !isDirectory(input) {
				paths = append(paths, input)
				continue
			}

			files, err := getAllDirectoryFiles(input)
			if err != nil {
				continue
			}
			paths = append(paths, files...)
		} else if isValidURL(input) {
			paths = append(paths, input)
		}
	}
	return paths
}

func readInputParams() []string {
	fmt.Println(inputMessage)
	var inputs []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs = strings.Fields(scanner.Text())
		if len(inputs) != 0 {
			break
		} else {
			fmt.Println(inputMessage)
		}
	}
	return inputs
}

func getAllDirectoryFiles(directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(
		directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
		})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func isValidURL(myURL string) bool {
	u, err := url.Parse(myURL)

	switch {
	case err != nil, u.Scheme == "", u.Host == "", u.Scheme != "http" && u.Scheme != "https":
		return false
	}

	return true
}

func isValidFileOrDirectoryPath(myPath string) bool {
	if _, err := os.Stat(myPath); err != nil {
		return false
	}
	return true
}

func isDirectory(myPath string) bool {
	fileInfo, err := os.Stat(myPath)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
