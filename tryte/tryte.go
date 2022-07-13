package tryte

import (
	"fmt"

	"github.com/thzoid/trivial/trit"
)

type Tryte [3]byte

func (t Tryte) String() string {
	str := make([]rune, TRYTE_TRIT)
	for i := 0; i < TRYTE_TRIT; i++ {
		str[i] = trit.TritToChar(trit.Trit((t[ByteOfTrit(uint64(i))] >> TritOffset(uint64(i)) & 0b11)))
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
