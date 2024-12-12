package zkdcap

import (
	"encoding/binary"
	"time"
)

const BONSAI_IMAGE_ID = "83613a8beec226d1f29714530f1df791fa16c2c4dfcf22c50ab7edac59ca637f"

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
