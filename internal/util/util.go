package util

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func GetMD5Hash(text string, len int) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:len])
}
