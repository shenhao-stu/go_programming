package routine

import (
	"fmt"
	"testing"
)

func inputNum(inputChan chan int, num int) {
	for i := 0; i < num; i++ {
		inputChan <- i
	}
	close(inputChan)
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	if num <= 0 || num == 1 {
		return false
	} else {
		return true
	}
}

func primeNum(inputChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-inputChan
		if !ok {
			break
		}
		if isPrime(num) {
			primeChan <- num
		}
	}
	exitChan <- true
}

func TestCalPrime(t *testing.T) {
	var inputChan = make(chan int, 200)
	var primeChan = make(chan int, 100)
	var exitChan = make(chan bool, 4)

	go inputNum(inputChan, 8000)

	for i := 0; i < 4; i++ {
		go primeNum(inputChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)

	}()

	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("Prime number is: %v \n", v)
	}

}
