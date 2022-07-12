package main

import (
	"fmt"

	"github.com/thzoid/trickster/tryte"
)

func main() {
	fmt.Println(tryte.Tryte{0b00011001}) // ------0+0
	fmt.Println(tryte.MustRead("-0+0"))  // ------0+0
}
