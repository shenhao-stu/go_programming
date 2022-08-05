package main

import "fmt"

func main() {
	// new 主要用来分配值类型内存，new函数的实参是一个类型，返回值是对应类型的指针 num: *int
	num := new(int)
	// type: *int, value: 0xc0000ac008, address: 0xc0000a6018, ptr_value: 0
	fmt.Printf("type: %T, value: %v, address: %v, ptr_value: %v\n", num, num, &num, *num)

	// make 主要用来分配引用类型
	// make底层创建一个数组，对外不可见，所以不可以直接操作这个数组，要通过slice去间接访问各个元素，不可以直接对数组进行操作
	// make(类型, 长度, 容量)
	slice := make([]int, 4, 20)        // 初始化一个切片，其中，切片类型为[]int，切片长度为4，切片容量为20
	fmt.Println(slice)                 // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println("切片的长度: ", len(slice)) // 4
	fmt.Println("切片的容量: ", cap(slice)) // 20
	slice[0] = 1
	slice[1] = 2
	fmt.Println(slice) // [1 2 0 0 0 0 0 0 0 0]

	slice2 := []int{1, 4, 7}            // slice
	arr := [3]int{1, 4, 7}              // array
	fmt.Println(slice2)                 // [1 4 7]
	fmt.Println(arr)                    // [1 4 7]
	fmt.Println("切片的长度: ", len(slice2)) // 4
	fmt.Println("切片的容量: ", cap(slice2)) // 20
}
