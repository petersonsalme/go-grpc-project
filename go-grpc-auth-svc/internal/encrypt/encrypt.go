package encrypt

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Password(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return base64.StdEncoding.EncodeToString(bytes)
}

func IsEquals(password, hash string) bool {
	decoded, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		log.Println("err: ", err)
		return false
	}

	return bcrypt.CompareHashAndPassword(decoded, []byte(password)) == nil
}
