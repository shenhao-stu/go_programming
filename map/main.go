package main

import "fmt"

/*
1.map集合在使用前一定要make
2.map的key-value是无序的
3.key是不可以重复的，如果遇到重复，后一个value会覆盖前一个value
4.value可以重复的
5.key通常为int、string，value通常为数字、string、map、结构体
6.make函数的第二个参数size可以省略，默认就分配一个内存
*/

func main() {
	// 方式1
	var a map[int]string = make(map[int]string, 10) // map 可以存10个键值对
	// 只声明map，内存是没有分配空间
	// 必须通过make函数进行初始化，才会分配空间
	a[1] = "a"
	a[2] = "b"
	a[3] = "c"
	fmt.Println(a)

	// 方式2
	b := make(map[int]string)
	b[1] = "a"
	b[2] = "b"
	b[3] = "c"
	fmt.Println(b)

	// 方式3
	c := map[int]string{1: "a", 2: "b", 3: "c"}
	c[4] = "d"
	fmt.Println(c)

	// 增删改查
	// 增改 b[key] = value
	// 删除 delete(map, key)，如果key存在，就删除该key-value，如果key不存在，不做任何操作，也不会报错
	delete(c, 1)
	fmt.Println(c)
	// 清空 map
	// 1.遍历逐个删除
	// 2.make一个新的
	// 查找
	// value 为返回的value，bool为是否返回
	value, flag := c[2]
	fmt.Println(value, flag)

	// 遍历
	for key, value := range c {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	// 扩展
	d := make(map[string]map[int]string)
	d["class1"] = make(map[int]string)
	d["class1"][1] = "a"
	d["class1"][2] = "b"
	d["class1"][3] = "c"
	d["class2"] = make(map[int]string)
	d["class2"][1] = "a"
	d["class2"][2] = "b"
	d["class2"][3] = "c"
	fmt.Println(d)
	for k, v := range d {
		fmt.Printf("key: %v\n", k)
		for key2, value2 := range v {
			fmt.Printf("key2: %v, value2: %v\n", key2, value2)
		}
	}
}
