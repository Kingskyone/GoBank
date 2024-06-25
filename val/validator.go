package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString                // 匹配用户名的正则表达式函数
	isValidFullName = regexp.MustCompile(`^[a-zA-Z0-9u4e00-u9fa5\s]+$`).MatchString // 匹配姓名的正则表达式函数
)

// ValidateString 验证字符串是否满足长度约束
func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("字符串长度不匹配")
	}
	return nil
}

// ValidateUsername 验证用户名
func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("用户名包含不允许字符")
	}
	return nil
}

// ValidatePassword 验证密码
func ValidatePassword(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}
	return nil
}

// ValidateEmail 验证Email
func ValidateEmail(value string) error {
	var err error
	if err = ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err = mail.ParseAddress(value); err != nil {
		return fmt.Errorf("不是有效的Email地址")
	}
	return nil
}

// ValidateFullName 验证姓名
func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("姓名包含不允许字符")
	}
	return nil
}
