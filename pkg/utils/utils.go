package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// MD5 计算字符串的MD5哈希值
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 计算字符串的SHA256哈希值
func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// RandomString 生成指定长度的随机字符串（包含大小写字母和数字）
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomNumber 生成指定长度的随机数字字符串
func RandomNumber(length int) string {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// CamelToSnake 将驼峰命名转换为蛇形命名（如: UserName -> user_name）
func CamelToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// SnakeToCamel 将蛇形命名转换为驼峰命名（如: user_name -> userName）
func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(string(parts[i][0])) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// Contains 检查字符串切片中是否包含指定元素
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// RemoveDuplicate 移除字符串切片中的重复元素
func RemoveDuplicate(slice []string) []string {
	keys := make(map[string]bool)
	result := []string{}
	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}

// FormatTime 按指定格式格式化时间
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseTime 按指定格式解析时间字符串
func ParseTime(s, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}

// InTimeRange 检查指定时间是否在时间范围内
func InTimeRange(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	return !check.Before(end) || !check.After(start)
}

// Paginate 计算分页偏移量和限制数
func Paginate(page, pageSize int) (offset, limit int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset = (page - 1) * pageSize
	limit = pageSize
	return
}

// MaskPhone 手机号脱敏（如: 13812345678 -> 138****5678）
func MaskPhone(phone string) string {
	if len(phone) != 11 {
		return phone
	}
	return phone[:3] + "****" + phone[7:]
}

// MaskEmail 邮箱脱敏（如: example@test.com -> e***@test.com）
func MaskEmail(email string) string {
	at := strings.Index(email, "@")
	if at <= 1 {
		return email
	}
	return email[:1] + "***" + email[at:]
}

// If 三元运算符，根据条件返回对应的值
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Ternary 三元运算符，If的别名
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	return If(condition, trueVal, falseVal)
}

// Must 必须成功，如果err不为nil则panic
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Default 如果值为零值则返回默认值
func Default[T comparable](val T, defaultVal T) T {
	var zero T
	if val == zero {
		return defaultVal
	}
	return val
}

// Chunk 将切片分割成指定大小的块
func Chunk[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[0:size:size])
	}
	if len(slice) > 0 {
		chunks = append(chunks, slice)
	}
	return chunks
}

// Map 对切片中的每个元素执行映射函数
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Filter 过滤切片中满足条件的元素
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 将切片归约为单个值
func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

// Find 查找切片中满足条件的第一个元素
func Find[T any](slice []T, f func(T) bool) (T, bool) {
	for _, v := range slice {
		if f(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Every 检查切片中所有元素是否都满足条件
func Every[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

// Some 检查切片中是否存在满足条件的元素
func Some[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

// Reverse 反转切片
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// Unique 移除切片中的重复元素
func Unique[T comparable](slice []T) []T {
	keys := make(map[T]bool)
	result := []T{}
	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Flatten 将二维切片扁平化为一维切片
func Flatten[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// GroupBy 按指定函数对切片元素进行分组
func GroupBy[T any, K comparable](slice []T, f func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Keys 获取map的所有键
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取map的所有值
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Entries 将map转换为键值对切片
func Entries[K comparable, V any](m map[K]V) []struct {
	Key   K
	Value V
} {
	entries := make([]struct {
		Key   K
		Value V
	}, 0, len(m))
	for k, v := range m {
		entries = append(entries, struct {
			Key   K
			Value V
		}{k, v})
	}
	return entries
}

// FromEntries 将键值对切片转换为map
func FromEntries[K comparable, V any](entries []struct {
	Key   K
	Value V
}) map[K]V {
	m := make(map[K]V)
	for _, e := range entries {
		m[e.Key] = e.Value
	}
	return m
}

// Pick 从map中选取指定的键
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Omit 从map中排除指定的键
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	omit := make(map[K]bool)
	for _, k := range keys {
		omit[k] = true
	}
	for k, v := range m {
		if !omit[k] {
			result[k] = v
		}
	}
	return result
}

// Merge 合并多个map
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// ToString 将任意类型转换为字符串
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// ToInt 将任意类型转换为int
func ToInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int64:
		return int(val), nil
	case float64:
		return int(val), nil
	case string:
		var i int
		_, err := fmt.Sscanf(val, "%d", &i)
		return i, err
	default:
		return 0, fmt.Errorf("cannot convert %T to int", v)
	}
}

// ToInt64 将任意类型转换为int64
func ToInt64(v interface{}) (int64, error) {
	switch val := v.(type) {
	case int:
		return int64(val), nil
	case int64:
		return val, nil
	case float64:
		return int64(val), nil
	case string:
		var i int64
		_, err := fmt.Sscanf(val, "%d", &i)
		return i, err
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", v)
	}
}

// ToFloat64 将任意类型转换为float64
func ToFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		var f float64
		_, err := fmt.Sscanf(val, "%f", &f)
		return f, err
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", v)
	}
}

// ToBool 将任意类型转换为bool
func ToBool(v interface{}) (bool, error) {
	switch val := v.(type) {
	case bool:
		return val, nil
	case string:
		return strings.ToLower(val) == "true" || val == "1", nil
	case int:
		return val != 0, nil
	default:
		return false, fmt.Errorf("cannot convert %T to bool", v)
	}
}
