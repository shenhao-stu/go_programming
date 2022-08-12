package demo

import (
	"fmt"
	"reflect"
	"testing"
)

// 使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值。
// 定义一个 Monster 结构体
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

// Monster 的方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 接收值，给 s 赋值
func (s Monster) Set(name string) Monster {
	s.Name = name
	return s
}

// 显示 s 的值
func (s Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}
func ReflectStruct(a interface{}) {
	// 获取 a 的 reflect.Type
	typ := reflect.TypeOf(a)
	// 获取a 的 reflect.Value
	val := reflect.ValueOf(a)
	// 获取 a 对应的类别
	kd := val.Kind()
	// 如果传入的不是 struct，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	// 获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num) //4
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		// 获取到 struct 标签, 注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果该字段有 tag 标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	// 获取该结构体有多少个方法
	// 如果传入的参数只是结构体的值传递，只计算接收者为结构体的方法
	// 如果传入的参数是结构体的指针传递，同时计算接收者为结构体或者结构体指针的方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// var params []reflect.Value
	// 方法的排序默认是按照函数名的排序（ASCII码）
	val.Method(1).Call(nil) // 获取到第二个方法。这里是 Print() 方法，并调用它。

	// 调用结构体的第1个方法Method(0)
	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) // 传入的参数是 []reflect.Value, 返回[] reflect.Value
	fmt.Println("res=", res[0].Int()) // 返回结果，返回的结果是 []reflect.Value

	// 调用结构体的第3个方法Method(2)
	var params2 []reflect.Value //声明了 []reflect.Value
	params2 = append(params2, reflect.ValueOf("marry"))
	newVal := val.Method(2).Call(params2)[0]
	fmt.Printf("newVal: %v\n", newVal)

}
func TestStruct(t *testing.T) {
	// 创建了一个 Monster 实例
	var a Monster = Monster{
		Name:  "tom",
		Age:   23,
		Score: 30.8,
	}
	// 将 Monster 实例传递给 TestStruct 函数
	ReflectStruct(a)
	// 由于传入的参数是结构体的指针类型
	rVal := reflect.ValueOf(&a)
	// 如果结构体的接收者是指针类型，或者值类型，此时都会被计入Methods
	fmt.Printf("rVal.NumMethod(): %v\n", rVal.NumMethod())

}
