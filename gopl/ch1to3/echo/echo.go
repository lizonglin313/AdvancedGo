package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var separator = flag.String("s", " ", "separator")

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

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *separator))
	if !*n {
		fmt.Println()
	}
}
