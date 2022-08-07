package main

import (
	"fmt"
	"io"
	"os"
)

// This exercise was completed only after looking at the 2 similar solutions below
// https://github.com/kdama/gopl/blob/master/ch07/ex02/main.go
// https://github.com/vinceyuan/gopl-solutions/blob/master/ch07/ex7.2/ex7.2.go

// ByteCounter structure simply wraps the interface to the passed object
// and the counter
type ByteCounter struct {
	writer io.Writer
	count  int64
}

// Write method expected by the io.Writer interface. Think of
// this method as an "intermediary" between the caller (e.g.
// Fprintf) and the actual type that implements io.Writer. So
// this method's purpose is only to get in the "call stack" to
// be able to modify the counter.
// Note that if the receiver argument is simply ByteCounter,
// then we're not modifying the count in any way, hence it needs
// to be a pointer
func (bc *ByteCounter) Write(p []byte) (int, error) {
	// Write to the underlying "medium" by calling its own Write
	// method, and pick up the number of bytes written
	n, err := bc.writer.Write(p)

	// Increase the counter. We'll add the number of bytes that
	// the "medium" reports as written, as opposed to the number
	// of bytes we're pushing to it. It's probably the correct way,
	// as the 2 values could be different (e.g. receive 10 bytes to
	// write but the underlying Write function will only manage to
	// write 6 bytes (perhaps due to a simple reason such as an error))
	bc.count += int64(n)

	return n, err
}

// CountingWriter only gets called once when the program starts. It's simply
// passed a "medium" that will be written to (an io.Writer), and it returns that
// "medium" along with a pointer to the counter of bytes written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	// will be 0 initially
	var count int64

	// Keep in mind that the struct here doesn't get in the way of the
	// ByteCounter type implementing the io.Writer interface, as the Write
	// method is defined against the ByteCounter type itself (the method
	// has a receiver argument of type ByteCounter)
	n :=  &ByteCounter{w, count}
	return n, &n.count
}

func main() {
	w, count := CountingWriter(os.Stdout)

	// An alternate way to get to the counter, via the ByteCounter object itself. See
	// https://stackoverflow.com/a/14289568/5853218 for the reason behind the syntax
	t := w.(*ByteCounter)

	// We'll print both the count obtained through the ByteCounter object, as well as
	// through the direct pointer, just to learn more

	// Initial check to see if we get 0
	fmt.Println(t.count, *count)

	fmt.Fprintf(w, "This is a line\n")
	fmt.Println(t.count, *count)
	fmt.Fprintf(w, "One more\n")
	fmt.Println(t.count, *count)
}
