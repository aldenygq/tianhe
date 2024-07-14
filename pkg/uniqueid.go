package pkg

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)
func GenerateUniqueID() (string,error)  {
	// 使用加密安全的随机数生成器生成字节切片
	byteSlice := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, byteSlice); err != nil {
		fmt.Printf("Failed to generate a random unique ID")
		return "",err
	}

	// 将字节切片转换为十六进制字符串
	hexStr := hex.EncodeToString(byteSlice)

	// 将十六进制字符串转换为小写
	return hexStr,nil
}