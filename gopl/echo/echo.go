package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Arg Index: %d   ", i)
		fmt.Println("Arg is:", os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)                              // 返回 string
	fmt.Println(strings.Join(os.Args[1:], " ")) // 返回 string
	fmt.Println(os.Args[1:])                    // 返回列表
}
