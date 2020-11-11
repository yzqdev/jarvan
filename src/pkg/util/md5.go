package util

import (
	"crypto/md5"
	"encoding/hex"
)
// md5 encode a string
func EncodeMD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}
