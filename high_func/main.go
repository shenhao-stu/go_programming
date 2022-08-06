package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(oper string) func(int, int) int {
	if oper == "+" {
		return add
	} else if oper == "-" {
		return sub
	} else {
		return nil
	}
}

func main() {
	fmt.Println(cal("+")(1, 2))
}
