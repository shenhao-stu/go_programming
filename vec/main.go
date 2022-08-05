package main

import "fmt"

func main(){
	// 定义一个二维数组
	var arr  = [...][3]int{{1,4,7},{2,5,8},{3,6,8}}
	fmt.Println(arr) // [[1 4 7] [2 5 8] [3 6 8]]
	for _, value := range arr {
		for _, v := range value {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}