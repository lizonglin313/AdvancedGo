package main

import (
	"flag"
	"fmt"
)

// Value is the interface to the value stored in a flag.
// type Value interface {
//     String() string
//     Set(string) error
// }

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscan(s, "%f%s", &value, &unit) // 从 s 中解析出参数
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	//
	//var w io.Writer
	//fmt.Printf("%T\n", w)
	//w = os.Stdout
	//fmt.Printf("%T\n", w)
	//w = new(bytes.Buffer)
	//fmt.Printf("%T\n", w)

	flag.Parse()
	fmt.Println(*temp)
}
