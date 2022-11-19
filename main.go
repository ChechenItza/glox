package main

import (
	"encoding/binary"
	c "interpreter-go/chunk"
	"interpreter-go/line"
	"interpreter-go/value"
)

func main() {
	va := value.ValueArray{}
	for i := 0; i < 512; i++ {
		va = append(va, value.Value(i))
	}

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 300)

	chunk := c.Chunk{
		Code: []byte{ c.OP_RETURN, c.OP_CONSTANT, 0, c.OP_CONSTANT_LONG, bs[2], bs[1], bs[0] },
		Constants: va,
		Lines: line.RleLines{ 3, 123, 4, 25 },
	}
	chunk.Disassemble("test chunk")
}

