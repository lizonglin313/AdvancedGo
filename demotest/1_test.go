package demotest

import (
	"fmt"
	"testing"
)

type Brid struct {
	Type string
}

func (b *Brid) BridType() string {
	return b.Type
}

type Brids interface {
	BridType() string
	BridName() string
	IsBig()
}

type Yanzi struct {
	Brid
	Name string
	Big  bool
}

func (y *Yanzi) BridName() string {
	return y.Name
}

func (y *Yanzi) IsBig() {
	if y.Big {
		fmt.Printf("brid %s is big!\n", y.BridName())
	} else {
		fmt.Printf("brid %s is not big!\n", y.BridName())
	}
}

type LaoYing struct {
	Brid
	Big  bool
	Name string
}

func (l *LaoYing) BridName() string {
	return l.Name
}

func (l *LaoYing) IsBig() {
	if l.Big {
		fmt.Printf("brid %s is big!\n", l.BridName())
	} else {
		fmt.Printf("brid %s is not big!\n", l.BridName())
	}
}

func BridsInfo(b Brids) {
	fmt.Printf("Look this is a %s, named %s\n", b.BridType(), b.BridName())

	b.IsBig()
}

func TestBrides(t *testing.T) {
	l := LaoYing{
		Brid: Brid{Type: "Ying"},
		Name: "YingYingYing",
		Big:  true,
	}

	y := Yanzi{
		Brid: Brid{
			Type: "YanZi",
		},
		Name: "YanZi",
		Big:  false,
	}

	BridsInfo(&l)
	BridsInfo(&y)

}
