package main

import "fmt"

// 通常来讲，go是值传递的
// 但是闭包除外，它是引用传递

func errorUsing() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func correctUsing() {
	for i := 0; i < 3; i++ {
		defer func(it int) {	// 每次迭代用新的局部变量记录i的值
			fmt.Println(it)
		}(i)
	}
}

func main() {
	errorUsing()
	correctUsing()
}
