package ternary

import "github.com/thzoid/trivial/tryte"

const (
	NOOP = iota
	COMP
	JMP
	STORE
	LOAD
	ADD
	HALT
)

type Program []tryte.Tryte
