package tryte

import (
	"fmt"

	"github.com/thzoid/trickster/tryte/std"
)

type Tryte [3]byte

func (t Tryte) String() string {
	str := make([]rune, std.TRYTE_TRIT)
	for i := 0; i < std.TRYTE_TRIT; i++ {
		str[i] = tritToChar(Trit((t[std.ByteOfTrit(uint64(i))] >> std.TritOffset(uint64(i)) & 0b11)))
	}
	return string(str)
}

func Read(s string) (Tryte, error) {
	if len(s) > std.TRYTE_TRIT {
		return Tryte{}, fmt.Errorf("string too long")
	}

	T := Tryte{}
	for i := 0; i < len(s); i++ {
		T[std.ByteOfTrit(uint64(std.TRYTE_TRIT-len(s)+i))] |= byte(charToTrit(s[i])) <<
			std.TritOffset(uint64(std.TRYTE_TRIT-len(s)+i))
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
