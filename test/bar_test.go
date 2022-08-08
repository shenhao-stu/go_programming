package test

import (
	"fmt"
	"testing"
)

var p *int

// 作用域，函数只能收到全局变量，另一个函数的局部变量不会影响bar()
// 函数的声明在外部
func bar() {
	//use p
	fmt.Println(*p)
}

func TestBar(t *testing.T) {
	a := 1
	p := &a
	bar()
	t.Log(*p)
}
