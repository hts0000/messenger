package password

import (
	"golang.org/x/crypto/argon2"
)

type Argon2Gen struct {
	defaultSalt []byte
}

func (a *Argon2Gen) EncryptPassword(password string, salt []byte) (encryptedPassword []byte, err error) {
	if len(salt) == 0 {
		salt = a.defaultSalt
	}
	encryptedPassword = argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return encryptedPassword, nil
}
