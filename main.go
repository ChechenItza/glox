package main

import (
	"encoding/binary"
)

func main() {
	va := ValueArray{}
	for i := 0; i < 512; i++ {
		va = append(va, Value(i))
	}

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 300)

	c := Chunk{
		Code:      []byte{OP_CONSTANT, 1, OP_CONSTANT_LONG, bs[1], bs[2], bs[3], OP_RETURN},
		Constants: va,
		Lines:     RleLines{2, 123, 5, 25},
	}
	Disassemble(&c, "test chunk")

	vm := NewVm()
	vm.Interpret(&c)
}
