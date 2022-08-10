package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

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

	// 4. rVal -> interface{}
	iV := rVal.Interface()
	// 如果不进行类型断言，无法执行iV.Name
	fmt.Printf("iV: %v, type: %T\n", iV, iV)

	// 5. interface{} -> Student
	switch v := iV.(type) {
	case Student:
		fmt.Printf("v.Name: %v, %T\n", v.Name, v)
	default:
		fmt.Printf("v: %v, %T\n", v, v)
	}
}

func main() {
	var stu Student = Student{Name: "tom", Age: 23}
	reflectTest(stu)
}
