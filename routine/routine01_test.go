package routine

import (
	"fmt"
	"testing"
	"time"
)

var (
	myMap = make(map[int]int)
)

func cal(n int) {
	count := 1
	for i := 1; i <= n; i++ {
		count *= i
	}

	myMap[n] = count
}

func TestCal(t *testing.T) {
	for i := 1; i <= 20; i++ {
		// 创建协程
		// 10个协程同时在myMap数据结构中写数据，因此会出现数据冲突
		go cal(i)
	}
	// 避免主进程直接遍历输出结束
	time.Sleep(time.Second * 1)
	for i, v := range myMap {
		fmt.Printf("%d: %d\n", i, v)
	}
}
