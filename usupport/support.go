package usupport

import(
	"strings"
	"math/rand"
	"crypto/sha256"
)


func Gen_Id() string{
	// randomly generate id
	// yoinked from: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	const n uint8 = 6

	id := make([]byte, n)
    // A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			id[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return strings.ToLower(string(id))
}

// move to utils
func Encrypt_Password(pass string) [32]byte{
	encoded_str := []byte(pass)
	encrypted_password := sha256.Sum256(encoded_str)
	return encrypted_password
}

// TODO: remember that the encrypted hash should match when Authing.


