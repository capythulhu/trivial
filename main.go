package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thzoid/trivial/ternary"
	"github.com/thzoid/trivial/tryte"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: ./main <memory> <tryte to get>")
		os.Exit(1)
	}

	providedMemory := os.Args[1]
	providedTryte, _ := strconv.ParseUint(os.Args[2], 10, 64)

	m := ternary.NewMemory(uint64((len(providedMemory) + tryte.TRYTE_TRIT - 1) / tryte.TRYTE_TRIT))
	for i := uint64(0); i < m.Size(); i++ {
		if i+1 < m.Size() {
			m.Set(i, tryte.MustRead(providedMemory[i*tryte.TRYTE_TRIT:i*tryte.TRYTE_TRIT+tryte.TRYTE_TRIT]))
		} else {
			m.Set(i, tryte.MustRead(providedMemory[i*tryte.TRYTE_TRIT:]))
		}
	}

	fmt.Println(m.Get(providedTryte))
}
