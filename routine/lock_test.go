package routine

import (
	"sync"
	"testing"
	"time"
)

// 对全局变量加锁
var (
	Map = make(map[int]int)
	// lock 是一个全局的互斥锁
	lock sync.Mutex
)

func lockcal(n int) {
	count := 1
	for i := 1; i <= n; i++ {
		count *= i
	}
	// 加锁
	// 10个协程同时在myMap数据结构中写数据，因此会出现数据冲突
	lock.Lock()
	Map[n] = count // concurrent map writes?
	// 解锁
	lock.Unlock()
}

func TestLockCal(t *testing.T) {
	for i := 1; i <= 20; i++ {
		// 创建协程
		go lockcal(i)
	}
	// 避免主进程直接遍历输出结束
	// 如果主线程休眠时间长了，会加长等待时间
	// 如果等待时间短了，可能还有协程处于工作状态，这时也会随着主线程的退出而销毁
	time.Sleep(time.Second * 1)

	// 避免读的时候，写操作还在进行，也需要加锁
	lock.Lock()
	for i, v := range Map {
		t.Logf("map[%d]: %d\n", i, v)
	}
	lock.Unlock()
}
