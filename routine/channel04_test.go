package routine

import (
	"fmt"
	"testing"
)

// 形参指定模式
// ch chan<- int, only write
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

// ch <-chan int, only read
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("v: %v\n", v)
	}
	var a struct{}
	exitChan <- a
}

func TestChannelType(t *testing.T) {
	var ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("ending......")

}
