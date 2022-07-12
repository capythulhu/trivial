package std

func ceil(a, b uint64) uint64 {
	return (a + b - 1) / b
}

// Bytes needed to store a number of trits
func BytesFromTrytes(T uint) (B uint) {
	return uint(ceil(uint64(T*TRYTE_TRIT), uint64(BYTE_TRIT)))
}

// Byte a trit is stored in
func ByteOfTrit(t uint64) (B uint64) {
	return t * TRIT_BIT / BYTE_BIT
}

// Byte a trit is stored in
func ByteOfTryte(T uint64) (B uint64) {
	return T*TRIT_BIT + T/BYTE_TRIT
}

// Offset of 0b11 inside a byte
func TritOffset(t uint64) (offset uint64) {
	return (BYTE_TRIT - 1 - t%BYTE_TRIT) * TRIT_BIT
}
