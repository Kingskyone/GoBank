package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 将密码进行bcrypt加密
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", password)
	}
	return string(hashedPassword), nil
}

// CheckPassword 检查密码是否错误
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
