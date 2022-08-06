package main

import (
	"errors"
	"fmt"
)

// recover() 函数用于捕获panic异常
// defer() 先进后出
func test() (err error) {
	/* // defer+recover来捕获panic异常
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("error: ", err)
		}
	}()*/

	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 使用errors.New()自定义错误
		return errors.New("can't divide by zero")
	} else {
		result := num1 / num2
		// 如果没有捕获err，程序出现错误后，程序被中断，不会继续执行
		fmt.Println(result)
		return nil
	}
}

func main() {
	err := test()
	if err != nil {
		fmt.Println("error: ", err)
		// 使用panic，程序被中断，不会继续执行
		panic(err)
	}
	fmt.Println("test() succeed")
	fmt.Println("run() succeed")
}
