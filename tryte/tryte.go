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

func Read(s string) (Tryte, error) {
	if len(s) > TRYTE_TRIT {
		return Tryte{}, fmt.Errorf("string too long")
	}

	T := Tryte{}
	for i := 0; i < len(s); i++ {
		T[ByteOfTrit(uint64(TRYTE_TRIT-len(s)+i))] |= byte(trit.CharToTrit(s[i])) <<
			TritOffset(uint64(TRYTE_TRIT-len(s)+i))
	}

	return T, nil
}

func MustRead(s string) Tryte {
	t, err := Read(s)
	if err != nil {
		panic(err)
	}
	return t
}
