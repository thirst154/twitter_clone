package libs

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

const SaltSize = 16

func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)

	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher.Write(passwordBytes)

	var hashedPasswwordBytes = sha512Hasher.Sum(nil)

	var hashedPasswordHex = hex.EncodeToString(hashedPasswwordBytes)

	return hashedPasswordHex
}

func CheckPassword(hashedPassword string, currPassword string, salt []byte) bool {
	var currentPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currentPasswordHash
}
