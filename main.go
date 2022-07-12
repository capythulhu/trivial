package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thzoid/trickster/memory"
	"github.com/thzoid/trickster/tryte"
	"github.com/thzoid/trickster/tryte/std"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: ./main <memory> <tryte to get>")
		os.Exit(1)
	}

	providedMemory := os.Args[1]
	providedTryte, _ := strconv.ParseUint(os.Args[2], 10, 64)

	m := memory.NewMemory(uint64((len(providedMemory) + std.TRYTE_TRIT - 1) / std.TRYTE_TRIT))
	for i := uint64(0); i < m.Size(); i++ {
		if i+1 < m.Size() {
			m.Set(i, tryte.MustRead(providedMemory[i*std.TRYTE_TRIT:i*std.TRYTE_TRIT+std.TRYTE_TRIT]))
		} else {
			m.Set(i, tryte.MustRead(providedMemory[i*std.TRYTE_TRIT:]))
		}
	}
	fmt.Println(m.Get(providedTryte))
}
