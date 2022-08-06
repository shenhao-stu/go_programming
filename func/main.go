package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	b := max
	fmt.Println(b(1, 2))
	fmt.Printf("%T", b)
}
