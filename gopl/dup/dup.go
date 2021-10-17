package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func printResult(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 统计终端重复
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 使用 ctrl C 终止程序输入并打印
		counts[input.Text()]++
	}

	printResult(counts)
}

// 统计文件中重复
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dupFileLines() {
	counts := make(map[string]int)
	files := os.Args[2:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupFileLines: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	printResult(counts)
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[2:] {
		data, err := ioutil.ReadFile(filename) // 返回字节切片
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	printResult(counts)
}

func main() {

	if os.Args[1] == "1" {
		dup1() // 以"流”模式读取输入，并根据需要拆分成多个行
	} else if os.Args[1] == "2" {
		dupFileLines() // 以"流”模式读取输入，并根据需要拆分成多个行
	} else if os.Args[1] == "3" {
		dup3() // 一口气把全部输入数据读到内存中，一次分割为多行
	}
}
