package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	return inorder(t)
}

func inorder(root *tree) string {
	if root == nil {
		return ""
	}

	return inorder(root.left) + fmt.Sprintf("%d ", root.value) + inorder(root.right)
}

func main() {
	// Build a generic binary tree
	one := tree{1, nil, nil}
	two := tree{2, &one, nil}
	five := tree{5, nil, nil}
	six := tree{6, nil, nil}
	four := tree{4, &five, &six}
	three := tree{3, &two, &four}

	// Use various ways to print the in-order sequence
	fmt.Println(inorder(&three))
	fmt.Printf("%s\n", three.String())
	fmt.Printf("%s\n", &three)
	fmt.Println(&three)
}
