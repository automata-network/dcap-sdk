package bonsai

func Groth16Encode(data []byte) []byte {
	out := make([]byte, 4+len(data))
	copy(out[:4], []byte{0xc1, 0x01, 0xb4, 0x2b})
	copy(out[4:], data)
	return out
}
