package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 返回值是一个结构体，类型是time.Time
	fmt.Println(now)  // 2022-08-05 11:35:37.711727 +0800 CST m=+0.000124131

	// 调用结构体的方法
	fmt.Printf("Year: %v \n", now.Year())        // 2022
	fmt.Printf("Month: %v \n", now.Month())      // August
	fmt.Printf("Month: %v \n", int(now.Month())) // 8
	fmt.Printf("Day: %v \n", now.Day())          // 5
	fmt.Printf("Hour: %v \n", now.Hour())        // 11
	fmt.Printf("Minute: %v \n", now.Minute())    // 35
	fmt.Printf("Second: %v \n", now.Second())    // 37

	// 日期的格式化
	// 将日期以年月日时分秒按照指定的格式输出为字符串
	// Sprintf可以得到字符串，以便后续使用
	datastr := fmt.Sprintln(now.Format("2006-01-02 15:04:05")) // 2022-08-05 11:35:37
	fmt.Println(datastr)                                       // 2022-08-05 11:35:37
}
