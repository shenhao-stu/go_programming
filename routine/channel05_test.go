package routine

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	// 使用select可以解决从管道取数据的阻塞问题

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i

	}

	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	// 问题 在实际开发中，我们可能不好确定什么时候关闭该管道
	// 可以使用select方式解决
label:
	for {
		select {
		// 如果intChan一直没有关闭，不会一直阻塞而导致deadlock
		// 会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("read intdata: %v\n", v)
		case v := <-strChan:
			fmt.Printf("read strdata: %v\n", v)
		default:
			fmt.Println("no data")
			break label
		}
	}
}
