package tryte

type Trit byte

func (t Trit) String() string {
	switch t {
	case 0b00:
		return "-"
	case 0b01:
		return "0"
	case 0b10:
		return "+"
	default:
		return "∄"
	}
}

func charToTrit(c byte) Trit {
	switch c {
	case '-':
		return 0b00
	case '0':
		return 0b01
	case '+':
		return 0b10
	default:
		return 0b11
	}
}

func Not(t Trit) (u Trit) {
	notX := ^t & 0b10
	y := t & 0b01
	u = notX ^ y<<1 | y
	return
}
