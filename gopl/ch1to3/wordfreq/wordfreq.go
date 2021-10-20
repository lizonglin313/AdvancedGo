package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitWord() {
	input := "foo   bar      baz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	wordMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordMap[input.Text()]++
	}
	for k, v := range wordMap {
		fmt.Printf("%s  %d\n", k, v)
	}

	splitWord()
}
