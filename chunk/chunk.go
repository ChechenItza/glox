package chunk

import (
	"fmt"
	"interpreter-go/line"
	"interpreter-go/value"
)

const (
	OP_RETURN = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
)

type Chunk struct {
	Code []byte
	Constants value.ValueArray
	Lines line.RleLines
}

func(chunk Chunk) Disassemble(name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(chunk.Code); {
		offset = chunk.disassembleInstruction(offset)
	}
}

func(chunk Chunk) disassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && chunk.Lines.Get(offset) == chunk.Lines.Get(offset - 1) {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", chunk.Lines.Get(offset))
	}

	instruction := chunk.Code[offset]
	switch instruction {
	case OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	case OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", chunk, offset)
	case OP_CONSTANT_LONG:
		return longConstantInstruction("OP_CONSTANT_LONG", chunk, offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}