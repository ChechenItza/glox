package interpreter

import (
	"encoding/binary"
)

const (
	OP_RETURN = iota
	OP_CONSTANT
	OP_CONSTANT_LONG
)

type Chunk struct {
	Code      []byte
	Constants ValueArray
	Lines     RleLines
}

func ConcatLongConstant(c ...byte) uint32 {
	test := []byte{0}
	test = append(test, c...)
	return binary.BigEndian.Uint32(test)
}
