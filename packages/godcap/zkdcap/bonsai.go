package zkdcap

import (
	"encoding/binary"
	"time"
)

const BONSAI_IMAGE_ID = "c2eafe1ba01610f3b71281f9dd3280b33d97370bb68d3ada2925d391be245e10"

func BonsaiGenerateInput(quote []byte, collateral *Collateral) []byte {
	currentTime := uint64(time.Now().Unix())

	var currentTimeBytes [8]byte
	binary.LittleEndian.PutUint64(currentTimeBytes[:], currentTime)

	collaterals := collateral.Encode()

	quoteLen := uint32(len(quote))
	collateralsLen := uint32(len(collaterals))

	totalLen := 8 + 4 + 4 + quoteLen + collateralsLen

	data := make([]byte, 0, totalLen)

	data = append(data, currentTimeBytes[:]...)

	var lenBuf [4]byte
	binary.LittleEndian.PutUint32(lenBuf[:], quoteLen)
	data = append(data, lenBuf[:]...)
	binary.LittleEndian.PutUint32(lenBuf[:], collateralsLen)
	data = append(data, lenBuf[:]...)

	data = append(data, quote...)
	data = append(data, collaterals...)

	return data
}
