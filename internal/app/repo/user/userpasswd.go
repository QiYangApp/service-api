package user

import (
	"context"
	"encoding/hex"
	"ent/models"
	"service-api/internal/app/repo"
	"service-api/internal/modules/auth/passwd/hash"
	"service-api/internal/modules/setting"
	"service-api/internal/modules/util"
)

func IsPasswdSet(user *models.User) bool {
	return len(user.Passwd) != 0
}

func ValidatePassword(user *models.User, passwd string) bool {
	return hash.Parse(user.PasswdHashAlgo).VerifyPassword(passwd, user.Passwd, user.PasswdSalt)
}

// SaltByteLength Note: As of the beginning of 2022, it is recommended to use at least
// 64 bits of salt, but NIST is already recommending to use to 128 bits.
// (16 bytes = 16 * 8 = 128 bits)
const SaltByteLength = 16

// GetUserSalt returns a random user salt token.
func GetUserSalt() (string, error) {
	rBytes, err := util.CryptoRandomBytes(SaltByteLength)
	if err != nil {
		return "", err
	}
	// Returns a 32 bytes long string.
	return hex.EncodeToString(rBytes), nil
}

func SetPassword(user *models.User, passwd string) (err error) {
	if user.PasswdSalt, err = GetUserSalt(); err != nil {
		return err
	}
	if user.Passwd, err = hash.Parse(setting.SecretSetting.PasswdHashAlgo).Hash(passwd, user.PasswdHashAlgo); err != nil {
		return err
	}

	user.PasswdHashAlgo = setting.SecretSetting.PasswdHashAlgo

	return nil
}

func UpdatePassword(ctx context.Context, user *models.User) (*models.User, error) {
	return repo.Client().User.
		UpdateOneID(user.ID).
		SetPasswd(user.Passwd).
		SetPasswdSalt(user.PasswdSalt).
		SetPasswdHashAlgo(user.PasswdHashAlgo).
		Save(ctx)
}
