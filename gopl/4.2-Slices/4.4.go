package main

import "fmt"

func main() {
	// rotate left by n elements

	fmt.Println(rotate([]int{1, 2, 3, 4, 5}, 2))
}

// Simply use the portion of the slice that will have to go
// in the beginning and then just stick to it the part that
// was previously in the front
func rotate(s []int, n int) []int {
	newSlice := s[n:]
	newSlice = append(newSlice, s[:n]...)

	return newSlice
}
