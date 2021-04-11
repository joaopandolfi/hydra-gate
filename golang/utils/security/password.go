package security

import (
	"hydra_gate/config"
	"hydra_gate/remotes/aes"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.Get().Server.Security.BcryptCost)
	return string(bytes), err
}

//CheckPasswordHash  compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateHash(val string) string {
	v, _ := aes.UnsecureEncrypt(config.Get().Server.Security.AESKey, val)
	return v
}
