package ternary

import "github.com/thzoid/trivial/tryte"

type CPU struct {
	memory Memory
	ip     tryte.Tryte
	acc    tryte.Tryte
}

func NewCPU(T uint64) (cpu CPU) {
	cpu.memory = NewMemory(T)
	return
}

func (cpu *CPU) Load(program []Instruction) {
	cpu.ip = tryte.UIntToUnb(0)
	cpu.acc = tryte.UIntToUnb(0)
	for i, j := range program {
		for k, T := range j.Rasterize() {
			cpu.memory.Set(uint64(i*tryte.TRYTE_TRIT+k), T)
		}
	}
}

func (cpu *CPU) Tick() {
	cpu.memory.Get(uint64(tryte.UnbToUInt(cpu.ip) * INSTRUCTION_SIZE))
}
