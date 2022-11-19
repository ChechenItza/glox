package debug

import (
	"fmt"
	"interpreter-go/chunk"
)

func Disassemble(c *chunk.Chunk, name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(c.Code); {
		offset = DisassembleInstruction(c, offset)
	}
}

func DisassembleInstruction(c *chunk.Chunk, offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.Lines.Get(offset) == c.Lines.Get(offset-1) {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.Lines.Get(offset))
	}

	instruction := c.Code[offset]
	switch instruction {
	case chunk.OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	case chunk.OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", c, offset)
	case chunk.OP_CONSTANT_LONG:
		return longConstantInstruction("OP_CONSTANT_LONG", c, offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}
