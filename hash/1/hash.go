package hash

import "encoding/binary"

func HashLen(input []byte) []byte {
	digest := make([]byte, 8)
	binary.BigEndian.PutUint64(digest, uint64(len(input)))
	return digest
}
