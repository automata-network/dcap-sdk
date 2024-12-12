package zkdcap

import (
	_ "embed"
	"encoding/binary"
	"time"
)

//go:embed riscv32im-succinct-zkvm-elf
var SUCCINCT_ZKVM_ELF []byte

func Sp1GenerateInput(quote []byte, collateral *Collateral) []byte {
	collateralBytes := collateral.Encode()
	currentTime := uint64(time.Now().Unix())
	totalLen := 8 + 4 + 4 + len(quote) + len(collateralBytes)
	data := make([]byte, 0, totalLen)

	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], currentTime)
	data = append(data, buf[:]...)

	binary.LittleEndian.PutUint32(buf[:4], uint32(len(quote)))
	data = append(data, buf[:4]...)

	binary.LittleEndian.PutUint32(buf[:4], uint32(len(collateralBytes)))
	data = append(data, buf[:4]...)

	data = append(data, quote...)

	data = append(data, collateralBytes...)
	return data
}
