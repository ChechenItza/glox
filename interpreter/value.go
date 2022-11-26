package interpreter

import "fmt"

type Value float32

func (v Value) Print() {
	fmt.Printf("%g", v)
}
