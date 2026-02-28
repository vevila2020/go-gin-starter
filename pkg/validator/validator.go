package validator

import (
	"regexp"
	"strings"
	"unicode"
)

// IsEmail 验证是否为有效的邮箱地址
func IsEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// IsPhone 验证是否为有效的中国大陆手机号
func IsPhone(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(pattern, phone)
	return matched
}

// IsIDCard 验证是否为有效的中国大陆身份证号（18位）
func IsIDCard(idCard string) bool {
	pattern := `^[1-9]\d{5}(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[\dXx]$`
	matched, _ := regexp.MatchString(pattern, idCard)
	return matched
}

// IsURL 验证是否为有效的URL
func IsURL(url string) bool {
	pattern := `^https?://[\w\-]+(\.[\w\-]+)+[/#?]?.*$`
	matched, _ := regexp.MatchString(pattern, url)
	return matched
}

// IsIP 验证是否为有效的IPv4地址
func IsIP(ip string) bool {
	pattern := `^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`
	matched, _ := regexp.MatchString(pattern, ip)
	return matched
}

// IsIPv6 验证是否为有效的IPv6地址
func IsIPv6(ip string) bool {
	pattern := `^([\da-fA-F]{1,4}:){7}[\da-fA-F]{1,4}$`
	matched, _ := regexp.MatchString(pattern, ip)
	return matched
}

// IsPort 验证是否为有效的端口号（1-65535）
func IsPort(port int) bool {
	return port >= 1 && port <= 65535
}

// IsUsername 验证是否为有效的用户名（3-20位字母数字下划线）
func IsUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	for _, r := range username {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return true
}

// IsPassword 验证是否为有效的密码（6-32位）
func IsPassword(password string) bool {
	return len(password) >= 6 && len(password) <= 32
}

// IsStrongPassword 验证是否为强密码（8-32位，包含大小写字母、数字和特殊字符）
func IsStrongPassword(password string) bool {
	if len(password) < 8 || len(password) > 32 {
		return false
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}

// IsAlpha 验证字符串是否只包含字母
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsAlphanumeric 验证字符串是否只包含字母和数字
func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsNumeric 验证字符串是否只包含数字
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsHex 验证字符串是否为有效的十六进制
func IsHex(s string) bool {
	s = strings.TrimPrefix(strings.ToLower(s), "0x")
	for _, r := range s {
		if !unicode.IsDigit(r) && !(r >= 'a' && r <= 'f') {
			return false
		}
	}
	return true
}

// IsBase64 验证字符串是否为有效的Base64编码
func IsBase64(s string) bool {
	pattern := `^[A-Za-z0-9+/]+={0,2}$`
	matched, _ := regexp.MatchString(pattern, s)
	return matched && len(s)%4 == 0
}

// IsJSON 验证字符串是否为有效的JSON格式
func IsJSON(s string) bool {
	s = strings.TrimSpace(s)
	return (strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")) ||
		(strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]"))
}

// IsUUID 验证字符串是否为有效的UUID
func IsUUID(uuid string) bool {
	pattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	matched, _ := regexp.MatchString(pattern, uuid)
	return matched
}

// IsDate 验证字符串是否为有效的日期格式（YYYY-MM-DD）
func IsDate(date string) bool {
	pattern := `^\d{4}-\d{2}-\d{2}$`
	matched, _ := regexp.MatchString(pattern, date)
	return matched
}

// IsDateTime 验证字符串是否为有效的日期时间格式（YYYY-MM-DD HH:mm:ss）
func IsDateTime(datetime string) bool {
	pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`
	matched, _ := regexp.MatchString(pattern, datetime)
	return matched
}

// IsTime 验证字符串是否为有效的时间格式（HH:mm:ss）
func IsTime(time string) bool {
	pattern := `^\d{2}:\d{2}:\d{2}$`
	matched, _ := regexp.MatchString(pattern, time)
	return matched
}

// IsCreditCard 验证字符串是否为有效的信用卡号（使用Luhn算法）
func IsCreditCard(card string) bool {
	card = strings.ReplaceAll(card, " ", "")
	card = strings.ReplaceAll(card, "-", "")

	if !IsNumeric(card) {
		return false
	}

	if len(card) < 13 || len(card) > 19 {
		return false
	}

	sum := 0
	alt := false
	for i := len(card) - 1; i >= 0; i-- {
		digit := int(card[i] - '0')
		if alt {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alt = !alt
	}

	return sum%10 == 0
}

// IsChinese 验证字符串是否全部为中文字符
func IsChinese(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// ContainsChinese 验证字符串是否包含中文字符
func ContainsChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// MinLength 验证字符串长度是否不小于指定值
func MinLength(s string, min int) bool {
	return len(s) >= min
}

// MaxLength 验证字符串长度是否不大于指定值
func MaxLength(s string, max int) bool {
	return len(s) <= max
}

// RangeLength 验证字符串长度是否在指定范围内
func RangeLength(s string, min, max int) bool {
	length := len(s)
	return length >= min && length <= max
}

// In 验证字符串是否在指定列表中
func In(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}

// NotIn 验证字符串是否不在指定列表中
func NotIn(s string, list []string) bool {
	return !In(s, list)
}

// Required 验证字符串是否非空（去除空格后）
func Required(s string) bool {
	return strings.TrimSpace(s) != ""
}

// Equals 验证两个字符串是否相等
func Equals(a, b string) bool {
	return a == b
}

// NotEquals 验证两个字符串是否不相等
func NotEquals(a, b string) bool {
	return a != b
}
