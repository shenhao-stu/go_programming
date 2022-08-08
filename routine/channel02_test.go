package routine

import (
	"fmt"
	"testing"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData: %v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			// readData 读取完数据后，即任务完成
			exitChan <- true
			close(exitChan)
			break
		}
		fmt.Printf("readData: %v\n", v)
	}
}

func TestChannel(t *testing.T) {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	// 阻塞等待
	for {
		v := <-exitChan
		if v {
			break
		}
	}
	// <-exitChan
}
