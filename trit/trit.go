package trit

type Trit byte

func (t Trit) String() string {
	return string(TritToChar(t))
}

func CharToTrit(c rune) Trit {
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

func TritToChar(t Trit) rune {
	switch t {
	case 0b00:
		return '-'
	case 0b01:
		return '0'
	case 0b10:
		return '+'
	default:
		return '∄'
	}
}
