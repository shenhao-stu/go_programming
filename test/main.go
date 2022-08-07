package main

import "fmt"

var p *int

// 作用域，函数只能收到全局变量，另一个函数的局部变量不会影响bar()
func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	a := 1
	p := &a
	bar()
	fmt.Println(*p)
}
