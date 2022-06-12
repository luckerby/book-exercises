package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "\x21\x20\xe2\xb5\xa3\x41"
	// Convert the string to a byte slice first, as strings
	// are immutable in Go and reversing in place won't work directly
	s := []byte(a)
	fmt.Printf("%s\n", s)
	fmt.Printf("% x\n", s)

	// First reverse the individual bytes for each Unicode char,
	// then reverse the whole byte slice
	ReverseIndividualBytesPerUnicodeChar(s)
	reverse(s)
	fmt.Printf("% x\n", s)
	fmt.Printf("%s\n", s)
}

// Decode the Unicode sets of bytes, and reverse them in place
func ReverseIndividualBytesPerUnicodeChar(s []byte) {
	//fmt.Printf("len(s)= %d\n", len(s))
	for i := 0; i < len(s); {
		//fmt.Printf("i= %d\n", i)
		_, size := utf8.DecodeRune(s[i:])
		//fmt.Printf("decoded %x (size=%d)\n", r, size)
		reverse(s[i : i+size])
		i += size
	}
}

// Blindly reverse the byte slice
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
