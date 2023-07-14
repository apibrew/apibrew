package util

import "golang.org/x/crypto/bcrypt"

func EncodeKey(key string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyKey(hash, key string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(key))
}
