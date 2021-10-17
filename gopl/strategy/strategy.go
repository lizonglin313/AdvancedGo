package main

import "fmt"

type IStrategy interface {
	do(int, int) int
}

type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

type sub struct{}

func (*sub) do(a, b int) int {
	return a - b
}

type Oprate struct {
	strategy IStrategy
}

func (op *Oprate) SetStrategy(is IStrategy) {
	op.strategy = is
}

func (op *Oprate) DoStrategy(a, b int) int {
	return op.strategy.do(a, b)
}

func main() {
	addS := add{}
	subS := sub{}

	op := &Oprate{}
	op.SetStrategy(&addS)
	fmt.Println(op.DoStrategy(1, 2))

	op.SetStrategy(&subS)
	fmt.Println(op.DoStrategy(2, 1))
}
