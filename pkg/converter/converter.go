package converter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// ToString 将任意类型转换为字符串
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// ToInt 将任意类型转换为int
func ToInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	case float32:
		return int(val), nil
	case float64:
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int", v)
	}
}

// ToInt64 将任意类型转换为int64
func ToInt64(v interface{}) (int64, error) {
	switch val := v.(type) {
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case string:
		return strconv.ParseInt(val, 10, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to int64", v)
	}
}

// ToFloat64 将任意类型转换为float64
func ToFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		return strconv.ParseFloat(val, 64)
	case bool:
		if val {
			return 1.0, nil
		}
		return 0.0, nil
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", v)
	}
}

// ToBool 将任意类型转换为bool
func ToBool(v interface{}) (bool, error) {
	switch val := v.(type) {
	case bool:
		return val, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return val != 0, nil
	case float32, float64:
		return val != 0.0, nil
	case string:
		return strconv.ParseBool(val)
	default:
		return false, fmt.Errorf("cannot convert %T to bool", v)
	}
}

// ToSlice 将任意类型转换为切片
func ToSlice[T any](v interface{}) ([]T, error) {
	switch val := v.(type) {
	case []T:
		return val, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to []T", v)
	}
}

// ToMap 将任意类型转换为map
func ToMap[K comparable, V any](v interface{}) (map[K]V, error) {
	switch val := v.(type) {
	case map[K]V:
		return val, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to map[K]V", v)
	}
}

// ToJSON 将任意类型转换为JSON字符串
func ToJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToJSONIndent 将任意类型转换为带缩进的JSON字符串
func ToJSONIndent(v interface{}) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON 从JSON字符串解析到指定类型
func FromJSON(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// ToTimestamp 将时间转换为Unix时间戳（秒）
func ToTimestamp(t time.Time) int64 {
	return t.Unix()
}

// ToTimestampMilli 将时间转换为Unix时间戳（毫秒）
func ToTimestampMilli(t time.Time) int64 {
	return t.UnixMilli()
}

// FromTimestamp 将Unix时间戳（秒）转换为时间
func FromTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// FromTimestampMilli 将Unix时间戳（毫秒）转换为时间
func FromTimestampMilli(ts int64) time.Time {
	return time.UnixMilli(ts)
}

// ParseTime 按指定格式解析时间字符串
func ParseTime(s, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}

// FormatTime 按指定格式格式化时间
func FormatTime(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseDate 解析日期字符串（格式：2006-01-02）
func ParseDate(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// FormatDate 格式化日期（格式：2006-01-02）
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// ParseDateTime 解析日期时间字符串（格式：2006-01-02 15:04:05）
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s)
}

// FormatDateTime 格式化日期时间（格式：2006-01-02 15:04:05）
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// MustToInt 将任意类型转换为int，转换失败返回零值
func MustToInt(v interface{}) int {
	val, _ := ToInt(v)
	return val
}

// MustToInt64 将任意类型转换为int64，转换失败返回零值
func MustToInt64(v interface{}) int64 {
	val, _ := ToInt64(v)
	return val
}

// MustToFloat64 将任意类型转换为float64，转换失败返回零值
func MustToFloat64(v interface{}) float64 {
	val, _ := ToFloat64(v)
	return val
}

// MustToBool 将任意类型转换为bool，转换失败返回零值
func MustToBool(v interface{}) bool {
	val, _ := ToBool(v)
	return val
}

// MustToJSON 将任意类型转换为JSON字符串，转换失败返回空字符串
func MustToJSON(v interface{}) string {
	s, _ := ToJSON(v)
	return s
}

// DefaultInt 将任意类型转换为int，转换失败返回默认值
func DefaultInt(v interface{}, defaultValue int) int {
	val, err := ToInt(v)
	if err != nil {
		return defaultValue
	}
	return val
}

// DefaultInt64 将任意类型转换为int64，转换失败返回默认值
func DefaultInt64(v interface{}, defaultValue int64) int64 {
	val, err := ToInt64(v)
	if err != nil {
		return defaultValue
	}
	return val
}

// DefaultFloat64 将任意类型转换为float64，转换失败返回默认值
func DefaultFloat64(v interface{}, defaultValue float64) float64 {
	val, err := ToFloat64(v)
	if err != nil {
		return defaultValue
	}
	return val
}

// DefaultBool 将任意类型转换为bool，转换失败返回默认值
func DefaultBool(v interface{}, defaultValue bool) bool {
	val, err := ToBool(v)
	if err != nil {
		return defaultValue
	}
	return val
}

// DefaultString 将任意类型转换为字符串，值为nil时返回默认值
func DefaultString(v interface{}, defaultValue string) string {
	if v == nil {
		return defaultValue
	}
	return ToString(v)
}
