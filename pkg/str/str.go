package str

import (
	"strings"
	"unicode"
)

// IsEmpty 检查字符串是否为空
func IsEmpty(s string) bool {
	return s == ""
}

// IsNotEmpty 检查字符串是否非空
func IsNotEmpty(s string) bool {
	return s != ""
}

// IsBlank 检查字符串是否为空白（空或只有空白字符）
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank 检查字符串是否非空白
func IsNotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}

// Trim 去除字符串两端的空白字符
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// TrimLeft 去除字符串左端的空白字符
func TrimLeft(s string) string {
	return strings.TrimLeft(s, " \t\n\r")
}

// TrimRight 去除字符串右端的空白字符
func TrimRight(s string) string {
	return strings.TrimRight(s, " \t\n\r")
}

// ToUpper 将字符串转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 将字符串转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Capitalize 将字符串首字母大写
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Uncapitalize 将字符串首字母小写
func Uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

// SwapCase 交换字符串中字母的大小写
func SwapCase(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if unicode.IsUpper(r) {
			result[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			result[i] = unicode.ToUpper(r)
		} else {
			result[i] = r
		}
	}
	return string(result)
}

// Reverse 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Contains 检查字符串是否包含子串
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsIgnoreCase 检查字符串是否包含子串（忽略大小写）
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// StartsWith 检查字符串是否以指定前缀开头
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith 检查字符串是否以指定后缀结尾
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// IndexOf 查找子串第一次出现的位置
func IndexOf(s, substr string) int {
	return strings.Index(s, substr)
}

// LastIndexOf 查找子串最后一次出现的位置
func LastIndexOf(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// Replace 替换所有匹配的子串
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// ReplaceN 替换前n个匹配的子串
func ReplaceN(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Split 按分隔符分割字符串
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitN 按分隔符分割字符串，最多分割n次
func SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

// Join 用分隔符连接字符串切片
func Join(parts []string, sep string) string {
	return strings.Join(parts, sep)
}

// Repeat 重复字符串指定次数
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// PadLeft 左填充字符串到指定长度
func PadLeft(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padLen := length - len(s)
	return strings.Repeat(pad, padLen) + s
}

// PadRight 右填充字符串到指定长度
func PadRight(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padLen := length - len(s)
	return s + strings.Repeat(pad, padLen)
}

// Center 居中填充字符串到指定长度
func Center(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padLen := length - len(s)
	leftPad := padLen / 2
	rightPad := padLen - leftPad
	return strings.Repeat(pad, leftPad) + s + strings.Repeat(pad, rightPad)
}

// Truncate 截断字符串并添加后缀
func Truncate(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + suffix
}

// Ellipsis 截断字符串并添加省略号
func Ellipsis(s string, length int) string {
	return Truncate(s, length, "...")
}

// Left 获取字符串左边的n个字符
func Left(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length]
}

// Right 获取字符串右边的n个字符
func Right(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[len(s)-length:]
}

// Mid 获取字符串中间的字符
func Mid(s string, start, length int) string {
	if start >= len(s) {
		return ""
	}
	end := start + length
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

// Substring 获取子字符串
func Substring(s string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return ""
	}
	return s[start:end]
}

// Count 统计子串出现的次数
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// Fields 按空白字符分割字符串
func Fields(s string) []string {
	return strings.Fields(s)
}

// Equals 检查两个字符串是否相等
func Equals(a, b string) bool {
	return a == b
}

// EqualsIgnoreCase 检查两个字符串是否相等（忽略大小写）
func EqualsIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)
}

// Compare 比较两个字符串
func Compare(a, b string) int {
	return strings.Compare(a, b)
}

// IsAlpha 检查字符串是否只包含字母
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsNumeric 检查字符串是否只包含数字
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsAlphanumeric 检查字符串是否只包含字母和数字
func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsWhitespace 检查字符串是否只包含空白字符
func IsWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsASCII 检查字符串是否只包含ASCII字符
func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

// RemovePrefix 移除字符串的前缀
func RemovePrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// RemoveSuffix 移除字符串的后缀
func RemoveSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// RemoveAll 移除字符串中所有匹配的子串
func RemoveAll(s, substr string) string {
	return strings.ReplaceAll(s, substr, "")
}

// RemoveSpaces 移除字符串中的所有空格
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// RemoveWhitespace 移除字符串中的所有空白字符
func RemoveWhitespace(s string) string {
	var result strings.Builder
	for _, r := range s {
		if !unicode.IsSpace(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// DefaultIfEmpty 如果字符串为空则返回默认值
func DefaultIfEmpty(s, defaultValue string) string {
	if IsEmpty(s) {
		return defaultValue
	}
	return s
}

// DefaultIfBlank 如果字符串为空白则返回默认值
func DefaultIfBlank(s, defaultValue string) string {
	if IsBlank(s) {
		return defaultValue
	}
	return s
}

// Abbreviate 缩写字符串到指定最大长度
func Abbreviate(s string, maxWidth int) string {
	if len(s) <= maxWidth {
		return s
	}
	return s[:maxWidth-3] + "..."
}

// Wrap 用指定字符串包裹字符串
func Wrap(s, wrap string) string {
	return wrap + s + wrap
}

// Unwrap 移除字符串两端的包裹字符
func Unwrap(s, wrap string) string {
	if StartsWith(s, wrap) && EndsWith(s, wrap) {
		return s[len(wrap) : len(s)-len(wrap)]
	}
	return s
}

// Quote 给字符串添加双引号
func Quote(s string) string {
	return `"` + s + `"`
}

// Unquote 移除字符串两端的双引号
func Unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}
	return s
}

// Lines 按换行符分割字符串
func Lines(s string) []string {
	return strings.Split(s, "\n")
}

// Words 按空白字符分割字符串
func Words(s string) []string {
	return strings.Fields(s)
}

// CamelToSnake 将驼峰命名转换为蛇形命名（如: UserName -> user_name）
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
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

// KebabToCamel 将短横线命名转换为驼峰命名（如: user-name -> userName）
func KebabToCamel(s string) string {
	parts := strings.Split(s, "-")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(string(parts[i][0])) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// CamelToKebab 将驼峰命名转换为短横线命名（如: UserName -> user-name）
func CamelToKebab(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('-')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
