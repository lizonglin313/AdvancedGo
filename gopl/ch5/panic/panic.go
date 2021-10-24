package main

import "fmt"

func panic0(x int) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case x:
			fmt.Println(x + 3)
		default:
			panic(p)
		}
	}()

	if x == 0 {
		panic(x)
	}
}

func main() {
	panic0(0)
}
