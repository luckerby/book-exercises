package main

func main() {
	a := [...]int{1, 2, 3, 4}
	reversePtr(&a)
}

// Note that the function stays the same as the one in the book,
// given the format for accessing via the pointer which is identical
// to the one used for a regular array
func reversePtr(s *[4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
