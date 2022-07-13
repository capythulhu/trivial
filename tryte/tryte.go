package tryte

import (
	"fmt"

	"github.com/thzoid/trivial/trit"
)

// Tryte data type
type Tryte [3]byte

// Convert tryte to string
func (T Tryte) String() string {
	str := make([]rune, TRYTE_TRIT)
	for i := uint64(0); i < TRYTE_TRIT; i++ {
		str[i] = trit.TritToChar(trit.Trit((T[ByteOfTrit(i)] >> TritOffset(i) & 0b11)))
	}
	return string(str)
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
		T[ByteOfTrit(uint64(TRYTE_TRIT-len(s)+i))] |= byte(trit.CharToTrit(r)) <<
			TritOffset(uint64(TRYTE_TRIT-len(s)+i))
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

// Map trits in tryte
func (T *Tryte) Map(callback func(trit.Trit) trit.Trit) {
	for i := uint64(0); i < TRYTE_TRIT; i++ {
		T.Set(i, callback(T.Get(i)))
	}
}
