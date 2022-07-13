package trit

// NOT gate
// xy | z
// 00 | 2
// 01 | 1
// 11 | X
// 10 | 0
//
// This implementation uses the following equation:
// x' = ~x ^ y; y' = y
func Not(t Trit) (u Trit) {
	notX := ^t & 0b10
	y := t & 0b01
	u = notX ^ y<<1 | y
	return
}

// INC gate
// xy | z
// 00 | 1
// 01 | 2
// 11 | X
// 10 | 0
//
// This implementation uses the following equation:
// x' = y; y' = ~x ^ y
func Inc(t Trit) (u Trit) {
	notX := ^t & 0b10
	y := t & 0b01
	u = y<<1 | notX>>1 ^ y
	return
}

// DEC gate
// xy | z
// 00 | 2
// 01 | 0
// 11 | X
// 10 | 1
func Dec(t Trit) (u Trit) {
	x := t & 0b10
	notY := ^t & 0b01
	u = x ^ notY<<1 | x>>1
	return
}

// ISF gate
// xy | z
// 00 | 2
// 01 | 0
// 11 | X
// 10 | 0
//
// This implementation uses the following equation:
// x' = ~x ^ y; y' = 0
func Isf(t Trit) (u Trit) {
	notX := ^t & 0b10
	y := t & 0b01
	u = notX ^ y<<1
	return
}

// ISU gate
// xy | z
// 00 | 0
// 01 | 2
// 11 | X
// 10 | 0
//
// This implementation uses the following equation:
// x' = y; y' = 0
func Isu(t Trit) (u Trit) {
	y := t & 0b01
	u = y << 1
	return
}

// IST gate
// xy | z
// 00 | 0
// 01 | 0
// 11 | X
// 10 | 2
//
// This implementation uses the following equation:
// x' = x; y' = 0
func Ist(t Trit) (u Trit) {
	x := t & 0b10
	u = x
	return
}

// CLD gate
// xy | z
// 00 | 0
// 01 | 1
// 11 | X
// 10 | 1
//
// This implementation uses the following equation:
// x' = 0; y' = x ^ y
func Cld(t Trit) (u Trit) {
	x := t & 0b10
	y := t & 0b01
	u = x>>1 ^ y
	return
}

// CLU gate
// xy | z
// 00 | 1
// 01 | 1
// 11 | X
// 10 | 2
//
// This implementation uses the following equation:
// x' = x; y' = ~x
func Clu(t Trit) (u Trit) {
	x := t & 0b10
	notX := ^t & 0b10
	u = x | notX>>1
	return
}

// MIN or AND gate
// zw  00 01 11 10
// xy  -----------
// 00 | 0  0  X  0
// 01 | 0  1  X  1
// 11 | X  X  X  X
// 10 | 0  1  X  2
//
// This implementation uses the following equation:
// x' = x & z; y' = ((x ^ y) & (z ^ w) & ~(x & z))
func And(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = x&z | ((x>>1 ^ y) & (z>>1 ^ w) &^ (x & z) >> 1)
	return
}

// MAX or OR gate
// zw  00 01 11 10
// xy  -----------
// 00 | 0  1  X  2
// 01 | 1  1  X  2
// 11 | X  X  X  X
// 10 | 2  2  x  2
//
// This implementation uses the following equation:
// x' = x | z; y' = (y | w) & ~(z ^ x)
func Or(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = x | z | ((y | w) & ^(z>>1 ^ x>>1))
	return
}

// NMIN or NAND
// zw  00 01 11 10
// xy  -----------
// 00 | 2  2  X  2
// 01 | 2  1  X  1
// 11 | X  X  X  X
// 10 | 2  1  X  0
//
// This implementation uses the following equation:
// x' = ~y' & ~(x & z); y' = y & (z | w) | x & w
func Nand(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	yL := (y&(z>>1|w) | x>>1&w) & 0b01
	notYL := ^yL & 0b01
	v = notYL<<1 & ^(x&z) | yL
	return
}

// NMAX or NOR
// zw  00 01 11 10
// xy  -----------
// 00 | 2  1  X  0
// 01 | 1  1  X  0
// 11 | X  X  X  X
// 10 | 0  0  X  0
//
// This implementation uses the following equation:
// x' = ~(x | y | z | w); y' = ~(x ^ z) & (y | w)
func Nor(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = ^(x|y<<1|z|w<<1)&0b10 | ^(x^z)>>1&(y|w)
	return
}

// XOR
// zw  00 01 11 10
// xy  -----------
// 00 | 0  1  X  2
// 01 | 1  1  X  1
// 11 | X  X  X  X
// 10 | 2  1  X  0
//
// This implementation uses the following equation:
// x' = x & (~z ^ w) | z & (~x ^ y); y' = w | y
func Xor(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = x&(^z^w<<1) | z&(^x^y<<1) | w | y
	return
}

// SUM
// zw  00 01 11 10
// xy  -----------
// 00 | 2  0  X  1
// 01 | 0  1  X  2
// 11 | X  X  X  X
// 10 | 1  2  X  0
//
// This implementation uses the following equation:
// x' = (~y ^ z) & (~x ^ w); y' = z & ~y & ~x| y & w | ~(z | w) & x
func Sum(t, u Trit) (v Trit) {
	x := t & 0b10
	notX := ^t & 0b10
	y := t & 0b01
	notY := ^t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = (notY<<1^z)&(^x^w<<1) | z>>1&notY&notX>>1 | y&w | ^(z>>1|w)&x>>1
	return
}

// CONSENSUS
// zw  00 01 11 10
// xy  -----------
// 00 | 0  1  X  1
// 01 | 1  1  X  1
// 11 | X  X  X  X
// 10 | 1  1  X  2
//
// This implementation uses the following equation:
// x' = x & z; y' = (x | y | z | w) & ~(x & z)
func Con(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = x&z | (x>>1|y|z>>1|w)&^(x&z)>>1
	return
}

// ANY
// zw  00 01 11 10
// xy  -----------
// 00 | 0  0  X  1
// 01 | 0  1  X  2
// 11 | X  X  X  X
// 10 | 1  2  X  2
//
// This implementation uses the following equation:
// x' = y & z | x & (z | w); y' = ~(x | y | w) & z | ~(y | z | w) & x | y & w
func Any(t, u Trit) (v Trit) {
	x := t & 0b10
	y := t & 0b01
	z := u & 0b10
	w := u & 0b01
	v = y<<1&z | x&(z|w<<1) | ^(x>>1|y|w>>1|z)&(y|z>>1|w)
	return
}
