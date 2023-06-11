package encryption

import "crypto/sha1"

func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func AesKeySecureRandom(keyword []byte) (key []byte) {
	data := []byte(keyword)

	hashs := SHA1(SHA1(data))
	key = hashs[0:16]

	return key
}
