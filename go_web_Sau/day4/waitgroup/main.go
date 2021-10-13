package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	// 结束时释放 Done
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup	// 声明一个等待组
	wg.Add(2)			// 增加两个信号
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()	// 阻塞 直到两个 Done 信号完成
	
}
