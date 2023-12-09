package authorize

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password, sing string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	password = fmt.Sprintf("%x-%s", p, sing)

	return fmt.Sprintf("%x", md5.Sum([]byte(password))), nil
}

func GenPasswordSing(ps, sing string) (string, error) {
	password := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%x-%s", ps, sing))))

	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", p), nil
}
