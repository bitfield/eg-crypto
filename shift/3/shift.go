package shift

func Encipher(plaintext []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b + 1
	}
	return ciphertext
}
