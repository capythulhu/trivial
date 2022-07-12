package tryte

import (
	"fmt"

	"github.com/thzoid/trickster/tryte/std"
)

type Tryte [3]byte

func (t Tryte) String() string {
	str := ""
	for i := 0; i < std.TRYTE_TRIT; i++ {
		str += Trit(t[std.ByteOfTrit(uint64(i))] >> std.TritOffset(uint64(i)) & 0b11).String()
	}
	return str
}

func FromString(s string) (Tryte, error) {
	if len(s) > std.TRYTE_TRIT {
		return Tryte{}, fmt.Errorf("string too long")
	}

	T := Tryte{}
	for i := 0; i < len(s); i++ {
		T[std.ByteOfTrit(uint64(i))] |= byte(charToTrit(s[i])) << std.TritOffset(uint64(i))
	}

	return T, nil
}
