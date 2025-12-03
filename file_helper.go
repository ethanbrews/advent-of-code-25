package main

import (
	"bufio"
	"bytes"
	"os"
)

// ReadLines reads a file and splits it by the given separator.
// If sep is "\n", it behaves like the original function.
func ReadLines(path string, sep string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ignore error for now

	var lines []string
	scanner := bufio.NewScanner(file)

	// Custom split function
	if sep != "\n" {
		s := []byte(sep)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if i := bytes.Index(data, s); i >= 0 {
				return i + len(s), data[:i], nil
			}
			if atEOF && len(data) > 0 {
				return len(data), data, nil
			}
			return 0, nil, nil
		})
	}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
