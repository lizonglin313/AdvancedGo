package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// 使用 ReadAll 写入内存 对大文件不适用
func fetchUsingMemory() {
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func fetchUsingIoCopy() {
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		f, err := os.Create("fetch.txt")
		defer f.Close()

		fileWriter := bufio.NewWriter(f)
		io.Copy(fileWriter, resp.Body)
		if err != nil {
			fmt.Printf("Copy to file: %v\n", err)
		}
	}
}

func main() {
	if os.Args[1] == "1" {
		fetchUsingMemory()
	} else {
		fetchUsingIoCopy()
	}
}
