package utils

import(
	"fmt"
	"errors"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"crypto/sha256"
)


func gen_id() string{
	// randomly generate id
	// yoinked from: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	const n int = 6

	id := make([]byte, n)
    // A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	
	return strings.ToLower(string(id))
}

// move to utils
func no_dupe_ids(new_user_id string) (string){

	user := employee_table.get(new_user_id)
	if user{
		new_id := _gen_id()    // gen new id & re-check for duplicate entry in db.
		_no_dupe_ids(new_id)
	}
	return new_user_id
}

// move to utils
func encrypt_password(pass string) string{
	encoded_str := []byte(pass)
	hash        := sha256.New()

	encrypted_password := hash.Write(encoded_str)
	return encrypted_password
}
