package hash

import "encoding/binary"

func SumHash(input []byte) []byte {
	digest := make([]byte, 8)
	var sum uint64
	for _, b := range input {
		sum += uint64(b)
	}
	binary.BigEndian.PutUint64(digest, sum)
	return digest
}
