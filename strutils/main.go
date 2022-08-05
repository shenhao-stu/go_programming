package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度，按字节进行统计
	str := "golang你好"     // 在golang中，汉字是utf-8字符集，一个汉字3个字节
	fmt.Println(len(str)) // 12

	// 对字符串进行遍历
	// 1.利用键值循环：for-range
	for i, v := range str {
		fmt.Printf("index: %d, value: %c \n", i, v)
	}

	// 2.利用切片 r:=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}

	// 字符串转整数
	num1, _ := strconv.Atoi("123")
	fmt.Println(num1)

	// 整数转字符串
	str1 := strconv.Itoa(123)
	fmt.Println(str1)

	// 统计一个字符串有几个指定的子串
	count := strings.Count(str, "你")
	fmt.Println(count) // 1

	// 不分大小写的字符串比较
	fmt.Println(strings.EqualFold("hello", "HELLO")) // true

	// 区分大小写进行字符串比较
	fmt.Println("hello" == "HELLO") // false

	// 返回子串在字符串中第一次出现的索引值，如果没有返回-1
	fmt.Println(strings.Index("hello", "e")) // 1

	// 字符串的替换
	// n可以指定替换的次数，如果n=-1，则替换所有匹配的子串
	fmt.Println(strings.Replace("hello", "e", "a", 1)) // hallo

	// 按照指定的字符为分割标识符，将字符串拆分成字符串数组
	fmt.Println(strings.Split("hello,world", ",")) // [hello world]

	// 将字符串的字母进行大小写转换
	fmt.Println(strings.ToUpper("hello")) // HELLO
	fmt.Println(strings.ToLower("HELLO")) // hello

	// 将字符串左右两边的空格去掉
	fmt.Println(strings.TrimSpace("  go and java  ")) // go and java

	// 将字符串左右两边的指定字符去掉
	fmt.Println(strings.Trim("~~go and java~", "~")) // go and java

	// 将字符串左边的指定字符去掉
	fmt.Println(strings.TrimLeft("~~go and java~", "~")) // go and java~

	// 将字符串右边的指定字符去掉
	fmt.Println(strings.TrimRight("~~go and java~", "~")) // ~~go and java

	// 判断字符串是否以指定的字符串开头
	fmt.Println(strings.HasPrefix("https://www.baidu.com", "https")) // true

	// 判断字符串是否以指定的字符串结尾
	fmt.Println(strings.HasSuffix("demo.png", "png")) // true
}
