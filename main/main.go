package main

import (
	"fmt"
	utils "go_programming/dbutils" // alias for package dbutils
)

var sum int = calNumber(10)

func init() {
	fmt.Println("func init() has been called")
}

func printStr(str string) {
	// special usage
	for i, v := range str {
		fmt.Printf("index: %v, value: %c\n", i, v)
	}
}

func calNumber(count int) (sum int) {
	fmt.Println("func calNumber() has been called")
	for i := 0; i <= count; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println("func main() has been called")
	fmt.Println(sum)
	printStr("hello world")
	utils.GetConnect("hello world")
}
