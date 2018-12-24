package util

import (
	"crypto/sha256"
  "encoding/hex"
)

/**
参数:sha256运算的内容
 */
func SHA256(content []byte) string{
	hash := sha256.New()
	hash.Write(content)
	//计算结果
	result := hash.Sum(nil)
	// 16进制的字符串
	str := hex.EncodeToString(result)
	return str
}
