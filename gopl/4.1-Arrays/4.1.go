package main

import (
	"crypto/sha256"
	"fmt"
)

// This code is implemented using a non-pre-computed function, unlike
// how the exercise reference hints
func main() {
	var d1 [32]byte = sha256.Sum256([]byte("test"))
	d2 := sha256.Sum256([]byte("teso"))
	fmt.Println(d1)
	fmt.Println(d2)

	//fmt.Println(CountDiffBits(8, 6))
	CountDiffBitsInHashes(d1, d2)

}

func CountDiffBitsInHashes(d1 [32]byte, d2 [32]byte) {
	for i := 0; i < 32; i++ {
		fmt.Println(CountDiffBits(d1[i], d2[i]))
	}
}

// CountDiffBits Both numbers are considered to be represented on the same number
// of bits
func CountDiffBits(n1 byte, n2 byte) int {
	count := 0
	t1 := n1
	t2 := n2
	for t1 > 0 || t2 > 0 {
		if t1%2 != t2%2 {
			count++
		}
		t1 = t1 / 2
		t2 = t2 / 2
	}
	return count

}

/*
func CountBits(n int) int {
	count := 0
	for i := n; i > 0; i = i / 2 {
		count += i % 2
	}
	return count
}
*/
