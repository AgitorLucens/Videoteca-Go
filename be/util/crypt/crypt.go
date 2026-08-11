package crypt

import "golang.org/x/crypto/bcrypt"

func PasswordHasher(password string) (string,error) {
	ph,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(ph), nil
}
func CheckPassword(password, hash string) (bool,error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false,err
	}
	return true, nil
}