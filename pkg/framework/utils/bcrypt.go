package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const bcryptDefaultCost = 31

func Md5(sing string) string {
	hash := md5.Sum([]byte(sing))

	// 将 md5 值转换为字符串
	return hex.EncodeToString(hash[:])
}

func GenerateFromPassword(password, sing string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcryptDefaultCost)
	if err != nil {
		return "", err
	}

	password = fmt.Sprintf("%x-%s", p, sing)

	return fmt.Sprintf("%x", md5.Sum([]byte(password))), nil
}

func GenPasswordSing(password, sing string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcryptDefaultCost)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", p), nil
}
