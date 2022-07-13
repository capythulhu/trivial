package tryte

const (
	// 1 char = 8 bits
	BYTE_BIT = 8
	// 1 trit = ceil(3, 2) = 2 bits
	TRIT_BIT = 2

	// 1 tryte = 9 trits
	TRYTE_TRIT = 9
	// 1 word = 3 trytes
	WORD_TRYTE = 3
	// 1 word = 27 trits
	WORD_TRIT = (WORD_TRYTE * TRYTE_TRIT)
	// 1 byte = 4 trits
	BYTE_TRIT = (BYTE_BIT / TRIT_BIT)

	// 1 tryte = ceil(TRYTE_TRIT, BYTE_TRIT) = 3 bytes
	TRYTE_BYTE = (TRYTE_TRIT + BYTE_TRIT - 1) / BYTE_TRIT
	// 1 word = ceil(WORD_TRIT, BYTE_TRIT) = 7 bytes
	WORD_BYTE = (WORD_TRIT + BYTE_TRIT - 1) / BYTE_TRIT
	// 1 word = ceil(WORD_TRIT*3, BYTE_TRIT) = 7 bytes
	TRIPLE_WORD_BYTE = (WORD_TRIT*3 + BYTE_TRIT - 1) / BYTE_TRIT

	// Heptavintimal = 27
	HEPTA_VINTIMAL = 27
	// 1 heptavintimal character = 3 trits
	HEPTA_VINTIMAL_TRIT = HEPTA_VINTIMAL / TRYTE_TRIT

	// Max value of a tryte = 3^9 - 1 = 19,682
	TRYTE_MAX = 19682
	// Max value of a balanced tryte = 3^9 / 2 - 1 = 9,840
	BTRYTE_MAX = 9840
	// Min value of a balanced tryte = 3^9 / 2 * -1 = -9,841
	BTRYTE_MIN = -9841
	// Max value of a tryte = 3^9 - 1 = 19,682
	WORD_MAX = 7625597484987

	// Amount of bytes for storing 3^81 - 1 =
	// 443,426,488,243,037,769,948,249,630,619,149,892,802 (136/8 = 17 bytes)
	TRIPLE_BYTES = 17
	// 1 Kitri (Kt) = 3^7 = 2,187
	KITRI = 2187
	// 1 Metri (Mt) = 3^7^2 = 2,187^2 = 4782,969
	METRI = 4782969
	// 1 Gitri (Gt) = 3^7^3 = 2,187^3 = 10,460,353,203
	GITRI = 10460353203
	// 1 Tetri (Tt) = 3^7^4 = 2,187^4 = 22,876,792,454,961
	TETRI = 22876792454961
)
