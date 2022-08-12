package main

import (
	"fmt"
	"unsafe"
)

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
	fmt.Printf("f type %T \n", e)
	fmt.Println(unsafe.Sizeof(e))

	// 字符型
	// 定义字符类型的数字，本质上就是一个整数，输出字符时，会将字符转换为对应的整数
	var ch byte = 'a'
	fmt.Println(ch) // 97
	// 汉字字符，底层对应的是Unicode码值
	var ch2 int = '中'
	fmt.Println(ch2)         // 20013
	fmt.Printf("%c \n", ch2) // 中

	// 转义字符
	var ch3 int = '\n'
	fmt.Println(ch3)         // 10
	fmt.Printf("%c \n", ch3) // 换行符

	// bool类型
	var b1 bool = true
	fmt.Println(b1) // true

	// 字符串类型
	var s1 string = "hello"
	fmt.Println(s1) // hello
	// 如果字符串中有特殊字符，字符串的表示形式用反引号表示
	var s2 string = `hello`
	fmt.Println(s2) // hello
	// 字符串不可变，s1[0]='t'会报错

	// 类型转换必须显式转换 float(g)
	// 基本数据类型和string类型之间的转换 fmt.Sprintf or strconv.FormatInt or strconv.FormatFloat or strconv.FormatBool
	var i int = 123
	var s3 string = fmt.Sprintf("%d", i) // 123
	// s
	// %d int %f float %t bool %c char
	fmt.Printf("%T, s3= %q \n", s3, s3) // string

	// 数组类型
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1) // [1 2 3]

	var a, b = 10, 20
	b, a = a, b
	fmt.Println(a, b)
}
