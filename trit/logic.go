package trit

func Not(t Trit) (u Trit) {
	notX := ^t & 0b10
	y := t & 0b01
	u = notX ^ y<<1 | y
	return
}
