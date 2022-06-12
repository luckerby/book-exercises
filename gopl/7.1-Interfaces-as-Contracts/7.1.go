package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// The solution here https://stackoverflow.com/a/61137191/5853218
// reduces the redunant code in both types
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	var err error
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	*c += WordCounter(count)

	return count, err
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	var err error
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	*c += LineCounter(count)

	return count, err
}

func main() {
	text := "This is a test sentence\nHello %s\nEnd\n"

	var w WordCounter
	var l LineCounter
	fmt.Printf(text, "Mihai")

	// Invoke the Write method on the objects
	fmt.Fprintf(&w, text, "Mihai")
	fmt.Fprintf(&l, text, "Mihai")
	fmt.Println(w) // prints the word count
	fmt.Println(l) // prints the line count
}
