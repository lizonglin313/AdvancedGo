package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {

	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("error of fetch %s: %s\n", url, err)
	}
	defer resp.Body.Close()

	numberOfBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		fmt.Printf("error of Copy :%s\n", err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%fs %7d  %s\n", secs, numberOfBytes, url)
}

func main() {
	ch := make(chan string)
	stime := time.Now()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Printf("%s", <-ch)
	}

	endtime := time.Since(stime).Seconds()
	fmt.Println(endtime, "s")

	time.Sleep(5 * time.Second)
}
