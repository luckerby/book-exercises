package main

import "fmt"

func main() {
	s := []string{"first", "two", "two", "two", "first", "first", "three"}
	s = eliminateAdjacentDuplicates(s)
	fmt.Println(s)
}

func eliminateAdjacentDuplicates(s []string) []string {
	lastString := ""
	i := 0
	for _, p := range s {
		if p != lastString {
			s[i] = p
			i++
			lastString = p
		}
	}

	return s[:i]
}
