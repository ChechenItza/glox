package debug

import (
	"fmt"
	"interpreter-go/chunk"
)

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, c *chunk.Chunk, offset int) int {
	i := c.Code[offset+1]

	fmt.Printf("%-16s %4d '", name, i)
	c.Constants[i].Print()
	fmt.Printf("'\n")

	return offset + 2
}

func longConstantInstruction(name string, c *chunk.Chunk, offset int) int {
	i := chunk.ConcatLongConstant(c.Code[offset+1 : offset+4]...)

	fmt.Printf("%-16s %4d '", name, i)
	c.Constants[i].Print()
	fmt.Printf("'\n")

	return offset + 4
}
