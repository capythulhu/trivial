package memory

import (
	"github.com/thzoid/trickster/tryte"
	"github.com/thzoid/trickster/tryte/std"
)

type Memory struct {
	buffer []byte
}

func NewMemory(capacity uint64) (m Memory) {
	m.buffer = make([]byte, std.BytesFromTrytes(capacity))
	return
}

func (m Memory) Size() (T uint64) {
	return uint64(len(m.buffer) * std.BYTE_TRIT / std.TRYTE_TRIT)
}

func (m Memory) Set(t uint64, value tryte.Tryte) {
	B := std.ByteOfTryte(t)
	offset := byte(B % std.TRYTE_TRIT)
	mask := byte(0xff >> offset)
	m.buffer[B+0] &= ^mask
	m.buffer[B+0] |= byte(value[0]) >> offset
	m.buffer[B+1] = byte(value[0]) << (std.BYTE_BIT - offset)
	m.buffer[B+1] |= byte(value[1]) >> offset
	mask >>= std.TRIT_BIT
	m.buffer[B+2] &= mask
	m.buffer[B+2] |= byte(value[1]) << (std.BYTE_BIT - offset)
	m.buffer[B+2] |= byte(value[2]&0b11000000) >> offset
}

func (m Memory) Get(t uint64) (value tryte.Tryte) {
	B := std.ByteOfTryte(t)
	offset := byte(B % std.TRYTE_TRIT)
	value[0] = m.buffer[B+0] << offset
	value[0] |= m.buffer[B+1] >> (std.BYTE_BIT - offset)
	value[1] = m.buffer[B+1] << offset
	value[1] |= m.buffer[B+2] >> (std.BYTE_BIT - offset)
	value[2] = m.buffer[B+2] << offset
	return
}
