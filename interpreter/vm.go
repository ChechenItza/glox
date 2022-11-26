package interpreter

import (
	"fmt"
)

type Result int

const (
	INTERPRET_OK Result = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type Vm struct {
	chunk *Chunk
	pc    int // TODO: use unsafe pointer arithmetics
}

func NewVm() Vm {
	return Vm{}
}

func (vm *Vm) Interpret(chunk *Chunk) Result {
	vm.chunk = chunk
	vm.pc = 0
	return vm.run()
}

func (vm *Vm) run() Result {
	for {
		DisassembleInstruction(vm.chunk, vm.pc)

		instruction := readByte(vm)
		switch instruction {
		case OP_RETURN:
			return INTERPRET_OK
		case OP_CONSTANT:
			constant := readConstant(vm)
			constant.Print()
			fmt.Printf("\n")
		case OP_CONSTANT_LONG:
			constant := readLongConstant(vm)
			constant.Print()
			fmt.Printf("\n")
		}
	}
}

func readByte(vm *Vm) byte {
	b := vm.chunk.Code[vm.pc]
	vm.pc++

	return b
}

func readConstant(vm *Vm) Value {
	op := readByte(vm)
	return vm.chunk.Constants[op]
}

func readLongConstant(vm *Vm) Value {
	op := ConcatLongConstant(readByte(vm), readByte(vm), readByte(vm))
	return vm.chunk.Constants[op]
}
