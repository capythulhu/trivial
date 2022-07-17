package ternary

import "github.com/thzoid/trivial/tryte"

const INSTRUCTION_SIZE = 2

type OpCode uint8

const (
	NOOP OpCode = iota
	LOAD
	STORE
	COMP
	JUMP
	CALL
	HALT
)

type Instruction struct {
	Op  OpCode
	Arg tryte.Tryte
}

func (i Instruction) Rasterize() [INSTRUCTION_SIZE]tryte.Tryte {
	return [INSTRUCTION_SIZE]tryte.Tryte{i.Arg, tryte.UIntToUnb(uint16(i.Op))}
}
