package main

import (
	"encoding/binary"
	"interpreter-go/interpreter"
)

func main() {
	va := make([]interpreter.Value, 512)
	for i := 0; i < 512; i++ {
		va[i] = interpreter.Value(i)
	}

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 300)

	c := interpreter.Chunk{
		Code:      []byte{interpreter.OP_CONSTANT, 1, interpreter.OP_CONSTANT_LONG, bs[1], bs[2], bs[3], interpreter.OP_RETURN},
		Constants: va,
		Lines:     interpreter.RleLines{2, 123, 5, 25},
	}
	interpreter.Disassemble(&c, "test chunk")

	vm := interpreter.NewVm()
	vm.Interpret(&c)
}
