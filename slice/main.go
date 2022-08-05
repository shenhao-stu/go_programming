package main

import "fmt"

/*
1.切片有三个字段的数据结构：一个是指向底层数组的指针，一个是切片的长度，一个是切片的容量。
2.切片定义后不可以直接使用 var slice []int ，需要让其引用到一个数组，或者make一个空间供切片来使用
3.切片使用不可以越界
4.简写方式：
var slice = arr[0:end] -> var slice = arr[:end]
var slice = arr[start:len(arr)] -> var slice = arr[start:]
var slice = arr[:len(arr)] -> var slice = arr[:]
5.切片还可以切片
6.切片可以动态增长
6.1 底层追加元素的时候对数组进行扩容，老数组扩容为心数组
6.2 创建一个新数组，将老数组中的4复制到新数组中，在新数组中追加9，50
6.3 slice4 底层数组的指向是新数组
6.4 往往我们在使用追加的时候其实想要做的效果给slice追加
6.5 底层的新数组不能直接维护，需要通过切片简介维护操作
7.copy(b,a) 将a中对应数组中元素内容复制到b中对应的数组中，互相不影响
*/
func main() {
	// 定义一个切片
	// 切片是对数组一个连续片段的引用
	var intarr = [6]int{3, 6, 9, 1, 4, 7}
	// [1:3] 表示从索引1开始，到索引3结束，不包括索引3
	var slice []int = intarr[1:3]
	fmt.Println(slice) //[6 9]
	// 切片元素的个数
	fmt.Println("slice的元素个数: ", len(slice)) // 2
	// 切片的容量，容量可以动态变化
	fmt.Println("slice的容量: ", cap(slice)) // 5

	// 引用类型
	slice[1] = 16
	fmt.Println(slice)  //[6 16]
	fmt.Println(intarr) //[3 6 16 1 4 7]

	// 定义切片
	slice2 := []int{1, 4, 7}
	for i, v := range slice2 {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// 切片的切片
	slice3 := slice2[1:2]
	fmt.Println(slice3) //[4]

	// 追加元素
	slice3 = append(slice3, 9, 50) // 因为指向新数组，所以对slice2不会有影响
	fmt.Println(slice3)            //[4 9 50]
	slice3[0] = 100
	fmt.Println(slice3) //[100 9 50]
	fmt.Println(slice2) //[4]

	// 追加切片
	slice3 = append(slice3, slice2...)
	fmt.Println(slice3) //[100 9 50 4]
}
