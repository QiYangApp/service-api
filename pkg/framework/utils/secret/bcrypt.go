package secret

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

const bcryptDefaultCost = 31

func Md5(sing string) string {
	hash := md5.Sum([]byte(sing))

	// 将 md5 值转换为字符串
	return hex.EncodeToString(hash[:])
}

func Sha3Sum512(sing string) string {
	hash := sha3.Sum512([]byte(sing))

	// 将 md5 值转换为字符串
	return hex.EncodeToString(hash[:])
}
func Sha1Sum(sing string) string {
	hash := sha1.Sum([]byte(sing))

	// 将 md5 值转换为字符串
	return hex.EncodeToString(hash[:])
}
