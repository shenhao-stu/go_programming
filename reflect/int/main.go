package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func reflectTest(b interface{}) {

	// 通过反射获取传入变量的type，kind，值
	// 1. reflect.Type 类型
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v, rType=%T\n", rType, rType)

	// 2. reflect.Value 类型
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v, rVal=%T\n", rVal, rVal)

	// 3. Kind 类型
	// rType.Kind() or rVal.Kind()
	fmt.Printf("kind = %v, kind = %T\n", rType.Kind(), rType.Kind())
	fmt.Printf("kind = %v, kind = %T\n", rVal.Kind(), rVal.Kind())

	n := 2 + rVal.Int()
	fmt.Printf("n: %v\n", n)

	// 4. rVal -> interface{}
	iV := rVal.Interface()

	// 5. interface{} -> int
	// num := iV.(int)
	switch num := iV.(type) {
	case MyInt:
		fmt.Printf("num: %v\n", num)
	case int:
		fmt.Printf("num: %v\n", num)
	}

}

func main() {
	var num MyInt = 10
	reflectTest(num)
}
