package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5 md5加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Md5Salt
func Md5Salt(s, salt string) string {
	return Md5(Md5(s) + salt)
}
