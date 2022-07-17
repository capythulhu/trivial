package tryte

import (
	"fmt"
	"math"

	"github.com/thzoid/trivial/trit"
)

// Tryte data type
type Tryte [TRYTE_BYTE]byte

// Convert tryte to string
func (T Tryte) String() string {
	str := make([]rune, TRYTE_TRIT)
	for i := uint64(0); i < TRYTE_TRIT; i++ {
		str[i] = trit.TritToChar(T.Get(i))
	}
	return string(str)
}

// Create unsigned tryte
func UIntToUnb(n uint16) (T Tryte) {
	l := uint64(TRYTE_TRIT - 1)
	for i := uint64(0); n > 0; i++ {
		T.Set(l-i, trit.Trit(n%3))
		n /= 3
	}
	return
}

// Create unsigned tryte
func UnbToUInt(T Tryte) (n uint16) {
	for i := uint64(0); i < TRYTE_TRIT; i++ {
		n += uint16(T.Get(i)) * uint16(math.Pow(3, float64(i)))
	}
	return
}

// Read balanced tryte from string
func Read(s string) (Tryte, error) {
	if len(s) > TRYTE_TRIT {
		return Tryte{}, fmt.Errorf("string too long")
	}

	T := Tryte{}
	for i, r := range s {
		if trit.CharToTrit(r) == 0b11 {
			return Tryte{}, fmt.Errorf("invalid character '%c'", r)
		}
		T.Set(uint64(TRYTE_TRIT-len(s)+i), trit.CharToTrit(r))
	}

	return T, nil
}

// Read balanced tryte from string (panics if string is invalid)
func MustRead(s string) Tryte {
	t, err := Read(s)
	if err != nil {
		panic(err)
	}

	return t
}

// Set trit in tryte
func (T *Tryte) Set(i uint64, t trit.Trit) {
	T[ByteOfTrit(i)] &= ^(0b11 << TritOffset(i))
	T[ByteOfTrit(i)] |= byte(t) << TritOffset(i)
}

// Get trit in tryte
func (T Tryte) Get(i uint64) (t trit.Trit) {
	return trit.Trit((T[ByteOfTrit(i)] >> TritOffset(i)) & 0b11)
}
