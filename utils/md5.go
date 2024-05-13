package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/exp/rand"
	"strings"
	"time"
)

func Md5(src []byte) string {
	m := md5.New()
	// 对源数据进行加密
	m.Write(src)
	// 将加密结果转换为十六进制字符串
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

/*
下面用于密码加密
*/

// GenerateSalt 为密码生成盐值
func GenerateSalt(length int) string {
	const alphanumeric = "abcdefghijklmnopqrstuvwxyz_0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, length)
	rand.Seed(uint64(time.Now().UnixNano()))
	for i := range bytes {
		bytes[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(bytes)
}

func PasswordMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
func PasswordMD5(str string) string {
	return strings.ToUpper(PasswordMd5(str))
}

// EncryptPassword 使用MD5和盐值加密密码
func EncryptPassword(password string, salt string) (encryptPassword string) {
	return PasswordMD5(password + salt)
}

// ValidPassword 检验密码
func ValidPassword(password, salt string, sqlPassword string) bool {
	return PasswordMD5(password+salt) == sqlPassword
}
