package pkg

import (
	"fmt"
	"regexp"
	"errors"
)

// 验证密码复杂度
func ValidatePassword(password string) error {
	// 至少8个字符，至少包含一个数字和一个特殊字符
	if len(password) < 8 {
        return fmt.Errorf("password len is < 9")
    }
    num := `[0-9]{1}`
    a_z := `[a-z]{1}`
    A_Z := `[A-Z]{1}`
    symbol := `[!@#~$%^&*()+|_]{1}`
    if b, _:= regexp.MatchString(num, password); !b {
        return errors.New("password need 0-9")
    }
    if b, _:= regexp.MatchString(a_z, password); !b {
        return errors.New("password need a-z")
    }
    if b,_ := regexp.MatchString(A_Z, password); !b {
        return errors.New("password need A-Z")
    }
    if b, _:= regexp.MatchString(symbol, password); !b {
        return errors.New("password need symbol")
    }
    return nil
}
