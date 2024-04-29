package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src []byte) string {
	m := md5.New()
	// 对源数据进行加密
	m.Write(src)
	// 将加密结果转换为十六进制字符串
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
