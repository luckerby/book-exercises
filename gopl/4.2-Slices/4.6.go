package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var input = "test ぁ\tあ\t\tぃ   p block   45"

	fmt.Printf("%s\n", input)
	fmt.Printf("% x\n", input)

	seq := []byte(input)
	ret := squashSpace(seq)
	fmt.Printf("% x\n", seq[:ret])
	fmt.Printf("%s\n", seq[:ret])
}

func squashSpace(pslice []byte) int { // we only return the length, since the underlying array is available by the caller anyway
	var newSize int // new size of the underlying array that should be used
	slice := pslice[:0]
	prevRune := rune(0)
	var i int
	for i = 0; i < len(pslice); {
		r, size := utf8.DecodeRune(pslice[i:])
		//fmt.Printf("%v\tIsSpace=%v\t%v\tdiff-from-prev=%v\n", r, unicode.IsSpace(r), i, r!=prevRune)
		if !unicode.IsSpace(r) {
			slice = append(slice, slice[i:i+size]...)
			newSize += size
		}

		if r != prevRune && unicode.IsSpace(r) {
			slice = append(slice, ' ') // despite using append, it's still in-place rewriting as the underlying array will be eventually smaller than the original one
			newSize++                  // add just 1 as the space we're adding only occupies one byte
		}
		i += size
		prevRune = r
		//fmt.Printf("% x\n", slice)
	}

	//return slice
	return newSize
}
