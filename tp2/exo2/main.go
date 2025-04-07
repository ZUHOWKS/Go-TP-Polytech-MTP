package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	N := 10
	channel := make(chan int, N)

	for i := 0; i <= N; i++ {
		wg.Add(1)
		go incrementValue(channel)
	}

	channel <- 0
	wg.Wait()
	fmt.Println("main() : ", <-channel)

}

func incrementValue(ch chan int) {
	defer wg.Done()
	n := <-ch
	fmt.Println("fonction incrementValue() :", n, "+ 1", "=", n+1)
	n = n + 1
	ch <- n
}
