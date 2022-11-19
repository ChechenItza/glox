package chunk

import (
	"encoding/binary"
	"fmt"
)

func simpleInstruction(s string, offset int) int {
	fmt.Printf("%s\n", s)
	return offset + 1
}

func constantInstruction(s string, c Chunk, offset int) int {
	i := c.Code[offset + 1]

	fmt.Printf("%-16s %4d '", s, i)
	c.Constants[i].Print()
	fmt.Printf("'\n")

	return offset + 2
}

func longConstantInstruction(s string, c Chunk, offset int) int {
	test := []byte{0}
	test = append(test, c.Code[offset + 1 : offset + 4]...)
	i := binary.BigEndian.Uint32(test)

	fmt.Printf("%-16s %4d '", s, i)
	c.Constants[i].Print()
	fmt.Printf("'\n")

	return offset + 4
}