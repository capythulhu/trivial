package ternary

import "github.com/thzoid/trivial/tryte"

type CPU struct {
	memory Memory
	ip     uint64 // TODO: make this a tryte
}

func NewCPU(T uint64) (cpu CPU) {
	cpu.memory = NewMemory(T)
	return
}

func (cpu *CPU) Load(program []tryte.Tryte) {
	cpu.ip = 0
	for i, t := range program {
		cpu.memory.Set(uint64(i), t)
	}
}

// func (cpu *CPU) Tick() {
// 	op := cpu.memory.Get(cpu.ip)
// 	cpu.ip++
// }
