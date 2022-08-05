package main

import "fmt"

/*
1.结构体是值类型
2.结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名称、个数和类型），s=Person(t)
3.结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是互相间可以强制转换 type Person1 = Person
var s1 Person
var s2 Person1
s1 = Person(s2)
*/
type Person struct {
	// 变量名大写，表示public
	Name string
	Age  int
}

func main() {
	// 默认值为{ 0}
	var obj Person = Person{}
	obj.Name = "张三"
	obj.Age = 18
	fmt.Println(obj) // {张三 18}

	// 初始化
	var obj1 Person = Person{"李四", 20}
	fmt.Println(obj1) // {李四 20}

	// 初始化指针
	var obj2 *Person = new(Person)
	// var obj2 *Person = &Person{}
	(*obj2).Name = "王五"
	(*obj2).Age = 21
	// 为了符合程序员的编程习惯，(*t).Name -> t.Name
	fmt.Println(*obj2) // {王五 21}

}
