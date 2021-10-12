package main

import (
	"fmt"
	"time"
)

func printNumber1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1()  {
	printNumber1()
	printLetters1()
}

func goPrint1() {
	go printNumber1()
	go printLetters1()
}

func printNumber2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func goPrint2() {
	go printNumber2()
	go printLetters2()
}

func main() {

}
