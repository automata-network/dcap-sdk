package bonsai

func Groth16Encode(data []byte) []byte {
	out := make([]byte, 4+len(data))
	copy(out[:4], []byte{0x50, 0xbd, 0x17, 0x69})
	copy(out[4:], data)
	return out
}
