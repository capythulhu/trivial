package tryte

import (
	"fmt"

	"github.com/thzoid/trickster/tryte/std"
)

type Tryte [3]byte

func (t Tryte) String() string {
	str := make([]rune, std.TRYTE_TRIT)
	for i := 0; i < std.TRYTE_TRIT; i++ {
		str[std.TRYTE_TRIT-1-i] = tritToChar(Trit(t[std.ByteOfTrit(uint64(i))] >> std.TritOffset(uint64(i)) & 0b11))
	}
	return string(str)
}

func Read(s string) (Tryte, error) {
	if len(s) > std.TRYTE_TRIT {
		return Tryte{}, fmt.Errorf("string too long")
	}

	T := Tryte{}
	for i := len(s) - 1; i >= 0; i-- {
		T[std.ByteOfTrit(uint64(i))] |= byte(charToTrit(s[len(s)-1-i])) << std.TritOffset(uint64(i))
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
