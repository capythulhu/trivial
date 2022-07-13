package ternary

import (
	"strconv"

	"github.com/thzoid/trivial/tryte"
)

type Memory struct {
	buffer []byte
}

func NewMemory(capacity uint64) (m Memory) {
	m.buffer = make([]byte, tryte.BytesFromTrytes(capacity))
	return
}

func (m Memory) Size() (T uint64) {
	return uint64(len(m.buffer) * tryte.BYTE_TRIT / tryte.TRYTE_TRIT)
}

func (m Memory) Set(T uint64, t tryte.Tryte) {
	B := tryte.ByteOfTryte(T)
	offset := byte(B % tryte.TRYTE_TRIT)
	mask := byte(0xff >> offset)
	m.buffer[B+0] &= ^mask
	m.buffer[B+0] |= byte(t[0]) >> offset
	m.buffer[B+1] = byte(t[0]) << (tryte.BYTE_BIT - offset)
	m.buffer[B+1] |= byte(t[1]) >> offset
	mask >>= tryte.TRIT_BIT
	m.buffer[B+2] &= mask
	m.buffer[B+2] |= byte(t[1]) << (tryte.BYTE_BIT - offset)
	m.buffer[B+2] |= byte(t[2]&0b11000000) >> offset
}

func (m Memory) Get(T uint64) (t tryte.Tryte) {
	B := tryte.ByteOfTryte(T)
	offset := byte(B % tryte.TRYTE_TRIT)
	t[0] = m.buffer[B+0] << offset
	t[0] |= m.buffer[B+1] >> (tryte.BYTE_BIT - offset)
	t[1] = m.buffer[B+1] << offset
	t[1] |= m.buffer[B+2] >> (tryte.BYTE_BIT - offset)
	t[2] = m.buffer[B+2] << offset
	return
}

func (m Memory) String() string {
	str := ""
	for _, b := range m.buffer {
		str += strconv.FormatInt(int64(b), 2) + " "
	}
	return str
}
