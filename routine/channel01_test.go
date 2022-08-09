package routine

/*
channel 本质就是一个数据结构-队列
channel 是引用类型 important
channel 必须初始化才能写入数据，即make后才能使用
intChan := make(chan int, capacity)
channel 只能存放指定的数据类型
在没有使用协程的情况下，channel 的数据放满后，就不能放入了，如果继续放入，会导致死锁，如果从channel取出数据，可以继续放入
在没有使用协程的情况下，如果我们的管道数据已经全部去除，再取就会报告 deadlock(不要用testing框架测试)
在使用协程的情况下，如果我们的管道数据已经全部去除，再取就会阻塞
*/

/*
### channel的关闭

使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据。并且取完所有数据后，会取默认值的数据

### channel的遍历

channel支持for-range的方式进行遍历，请注意两个细节

- 在遍历时，如果channel没有关闭，则会出现deadlock的错误
- 在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成后，就会退出遍历
*/

import (
	"fmt"
	"testing"
)

// 使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据
func TestCloseChannel(t *testing.T) {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // close
	// 再向channel写数据，会报错
	// intChan <- 300 // panic: send on closed channel
	// 当管道关闭后，仍可以读取数据
	fmt.Println(<-intChan) // 100
	fmt.Println(<-intChan) // 200
	fmt.Println(<-intChan) // 0
	fmt.Println(<-intChan) // 0

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// 遍历
	// 不能用len(intChan2)遍历，每次循环都会-1
	// 在遍历时，如果channel没有关闭，则会出现deadlock的错误
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}

}
