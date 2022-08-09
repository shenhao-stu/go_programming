package routine

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var lock sync.Mutex
	var wg sync.WaitGroup
	count := 0

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			count++
			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)

}
