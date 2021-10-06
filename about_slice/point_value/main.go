package main

import "fmt"

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type InterSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

func twice1(x *InterSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	twice(s)
	fmt.Println(s)

	h := &InterSliceHeader{Data: s, Len: len(s)}
	twice1(h)
	fmt.Println(h.Data)

}
