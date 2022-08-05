package main

import "fmt"

func test(arr [3]int) {
	// 值传递
	arr[0] = 50
}

func test_(arr *[3]int) {
	// 利用指针传递
	arr[0] = 50
}

func main() {
	// 定义一个数组
	var scores = [...]int{0, 0, 0, 0, 0}
	// var scores = [5]int{0, 0, 0, 0, 0}
	// var scores = [...]int{2:99, 0:89, 4:100}
	for i := 0; i < len(scores); i++ {
		fmt.Printf("Please input %d score:", i+1)
		fmt.Scanln(&scores[i])
	}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg := sum / len(scores)
	fmt.Printf("总分：%v, 平均分：%v\n", sum, avg)

	for i, v := range scores {
		fmt.Printf("%d: %d\n", i+1, v)
	}

	fmt.Println("----------------------------")
	var arr = [3]int{3, 6, 9}
	fmt.Println("arr: ", arr)
	test(arr)
	fmt.Println("After test() arr: ", arr)
	test_(&arr)
	fmt.Println("After test_() arr: ", arr)
}
