package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal Kind=%v\n", rVal.Kind()) // Kind=ptr

	// rVal.Elem() 返回reflect.Value持有的接口保管的值或持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Printf("num: %v\n", num)
}
