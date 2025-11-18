package util

import (
	"crypto/md5"
	"encoding/hex"
)

func SecPass(pass string) string {
	hash := md5.Sum([]byte(pass))
	return hex.EncodeToString(hash[:])
}
