package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"hash"

	"golang.org/x/crypto/bcrypt"
)

// MD5 计算字节数组的MD5哈希值
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// MD5String 计算字符串的MD5哈希值
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 计算字节数组的SHA1哈希值
func SHA1(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1String 计算字符串的SHA1哈希值
func SHA1String(s string) string {
	return SHA1([]byte(s))
}

// SHA256 计算字节数组的SHA256哈希值
func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256String 计算字符串的SHA256哈希值
func SHA256String(s string) string {
	return SHA256([]byte(s))
}

// SHA512 计算字节数组的SHA512哈希值
func SHA512(data []byte) string {
	h := sha512.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA512String 计算字符串的SHA512哈希值
func SHA512String(s string) string {
	return SHA512([]byte(s))
}

// HashPassword 使用bcrypt对密码进行哈希加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 验证密码是否与哈希值匹配
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Hash 哈希对象，支持链式调用
type Hash struct {
	h hash.Hash
}

// NewMD5 创建MD5哈希对象
func NewMD5() *Hash {
	return &Hash{h: md5.New()}
}

// NewSHA256 创建SHA256哈希对象
func NewSHA256() *Hash {
	return &Hash{h: sha256.New()}
}

// NewSHA512 创建SHA512哈希对象
func NewSHA512() *Hash {
	return &Hash{h: sha512.New()}
}

// Write 写入字节数据（链式调用）
func (h *Hash) Write(data []byte) *Hash {
	h.h.Write(data)
	return h
}

// WriteString 写入字符串数据（链式调用）
func (h *Hash) WriteString(s string) *Hash {
	h.h.Write([]byte(s))
	return h
}

// Sum 获取哈希值
func (h *Hash) Sum() string {
	return hex.EncodeToString(h.h.Sum(nil))
}

// Reset 重置哈希对象
func (h *Hash) Reset() {
	h.h.Reset()
}
