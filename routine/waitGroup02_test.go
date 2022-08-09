package routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncService(wg *sync.WaitGroup) chan string {
	retCh := make(chan string)
	(*wg).Add(1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
		(*wg).Done()
	}()
	return retCh
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}
func TestSyncChannel(t *testing.T) {
	var wg sync.WaitGroup
	retChan := AsyncService(&wg)
	otherTask()
	fmt.Println(<-retChan)
	wg.Wait() // 如果不使用WaitGroup，那么当<-retChan读取到数据，main就会结束，因此不会输出"service exited."
}
