package interpreter

import (
	"fmt"
)

func Disassemble(c *Chunk, name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(c.Code); {
		offset = DisassembleInstruction(c, offset)
	}
}

func DisassembleInstruction(c *Chunk, offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.Lines.Get(offset) == c.Lines.Get(offset-1) {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.Lines.Get(offset))
	}

	instruction := c.Code[offset]
	switch instruction {
	case OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	case OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", c, offset)
	case OP_CONSTANT_LONG:
		return longConstantInstruction("OP_CONSTANT_LONG", c, offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, c *Chunk, offset int) int {
	i := c.Code[offset+1]

	fmt.Printf("%-16s %4d '", name, i)
	c.Constants[i].Print()
	fmt.Printf("'\n")

	return offset + 2
}

func longConstantInstruction(name string, c *Chunk, offset int) int {
	i := ConcatLongConstant(c.Code[offset+1 : offset+4]...)

	fmt.Printf("%-16s %4d '", name, i)
	c.Constants[i].Print()
	fmt.Printf("'\n")

	return offset + 4
}
