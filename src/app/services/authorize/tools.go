package authorize

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const bcryptDefaultCost = 31

func Encrypt(password, sing string) (string, error) {
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
