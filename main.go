package main

import (
	"encoding/binary"
	c "interpreter-go/chunk"
	"interpreter-go/debug"
	"interpreter-go/line"
	"interpreter-go/value"
	vm2 "interpreter-go/vm"
)

func main() {
	va := value.ValueArray{}
	for i := 0; i < 512; i++ {
		va = append(va, value.Value(i))
	}

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 300)

	c := c.Chunk{
		Code:      []byte{c.OP_CONSTANT, 1, c.OP_CONSTANT_LONG, bs[1], bs[2], bs[3], c.OP_RETURN},
		Constants: va,
		Lines:     line.RleLines{2, 123, 5, 25},
	}
	debug.Disassemble(&c, "test chunk")

	vm := vm2.NewVm()
	vm.Interpret(&c)
}
