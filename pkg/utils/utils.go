package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum(nil))
}
