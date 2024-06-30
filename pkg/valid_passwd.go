package pkg

import (
	"fmt"
	"regexp"
)

// 验证密码复杂度
func ValidatePassword(password string) bool {
	// 至少8个字符，至少包含一个数字和一个特殊字符
	pattern := `^(?=.*[0-9])(?=.*[!@#$%^&*])(.{8,})$`
	matched, err := regexp.MatchString(pattern, password)
	if err != nil {
		fmt.Printf("正则表达式错误: %s\n", err)
		return false
	}
	return matched
}
