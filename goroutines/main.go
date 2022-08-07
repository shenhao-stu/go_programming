package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100) // 100ms
	}
}
func main() {
	go showMsg("java")
	showMsg("go")         // 在main协程中执行，如果它前面也添加go，程序没有输出
	fmt.Println("end...") // 主函数退出，程序就结束了
}
