package slice

// Contains 检查切片中是否包含指定元素
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// IndexOf 查找元素在切片中第一次出现的位置，不存在返回-1
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// LastIndexOf 查找元素在切片中最后一次出现的位置，不存在返回-1
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

// Remove 移除切片中所有匹配的元素
func Remove[T comparable](slice []T, item T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

// RemoveAt 移除切片中指定位置的元素
func RemoveAt[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// Insert 在切片指定位置插入元素
func Insert[T any](slice []T, index int, item T) []T {
	if index < 0 || index > len(slice) {
		return slice
	}
	result := make([]T, 0, len(slice)+1)
	result = append(result, slice[:index]...)
	result = append(result, item)
	result = append(result, slice[index:]...)
	return result
}

// Unique 移除切片中的重复元素
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Reverse 反转切片
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
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
	result := make([]T, 0)
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

// FindIndex 查找切片中满足条件的第一个元素的索引
func FindIndex[T any](slice []T, f func(T) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
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

// None 检查切片中是否所有元素都不满足条件
func None[T any](slice []T, f func(T) bool) bool {
	return !Some(slice, f)
}

// Count 统计元素在切片中出现的次数
func Count[T comparable](slice []T, item T) int {
	count := 0
	for _, v := range slice {
		if v == item {
			count++
		}
	}
	return count
}

// CountBy 统计切片中满足条件的元素个数
func CountBy[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, v := range slice {
		if f(v) {
			count++
		}
	}
	return count
}

// Chunk 将切片分割成指定大小的块
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	result := make([][]T, 0)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

// Flatten 将二维切片扁平化为一维切片
func Flatten[T any](slices [][]T) []T {
	result := make([]T, 0)
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

// Partition 将切片分成两部分：满足条件的和不满足条件的
func Partition[T any](slice []T, f func(T) bool) ([]T, []T) {
	passed := make([]T, 0)
	failed := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			passed = append(passed, v)
		} else {
			failed = append(failed, v)
		}
	}
	return passed, failed
}

// Union 求多个切片的并集
func Union[T comparable](slices ...[]T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0)
	for _, slice := range slices {
		for _, v := range slice {
			if !seen[v] {
				seen[v] = true
				result = append(result, v)
			}
		}
	}
	return result
}

// Intersection 求多个切片的交集
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	counts := make(map[T]int)
	for _, v := range slices[0] {
		counts[v] = 1
	}

	for _, slice := range slices[1:] {
		seen := make(map[T]bool)
		for _, v := range slice {
			if counts[v] > 0 && !seen[v] {
				counts[v]++
				seen[v] = true
			}
		}
	}

	result := make([]T, 0)
	for _, v := range slices[0] {
		if counts[v] == len(slices) {
			result = append(result, v)
			counts[v] = 0
		}
	}
	return result
}

// Difference 求两个切片的差集（在slice1但不在slice2中的元素）
func Difference[T comparable](slice1, slice2 []T) []T {
	seen := make(map[T]bool)
	for _, v := range slice2 {
		seen[v] = true
	}

	result := make([]T, 0)
	for _, v := range slice1 {
		if !seen[v] {
			result = append(result, v)
		}
	}
	return result
}

// Take 获取切片的前n个元素
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

// TakeRight 获取切片的后n个元素
func TakeRight[T any](slice []T, n int) []T {
	if n <= 0 {
		return nil
	}
	if n >= len(slice) {
		return slice
	}
	return slice[len(slice)-n:]
}

// Drop 丢弃切片的前n个元素
func Drop[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return nil
	}
	return slice[n:]
}

// DropRight 丢弃切片的后n个元素
func DropRight[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return nil
	}
	return slice[:len(slice)-n]
}

// First 获取切片的第一个元素
func First[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[0], true
}

// Last 获取切片的最后一个元素
func Last[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[len(slice)-1], true
}

// IsEmpty 检查切片是否为空
func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// Len 获取切片长度
func Len[T any](slice []T) int {
	return len(slice)
}

// Clone 克隆切片
func Clone[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

// Concat 连接多个切片
func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, slice := range slices {
		totalLen += len(slice)
	}
	result := make([]T, 0, totalLen)
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

// Fill 用指定值填充切片
func Fill[T any](slice []T, value T) []T {
	result := make([]T, len(slice))
	for i := range result {
		result[i] = value
	}
	return result
}

// Repeat 重复指定值生成切片
func Repeat[T any](value T, count int) []T {
	result := make([]T, count)
	for i := range result {
		result[i] = value
	}
	return result
}

// Slice 获取切片的子切片
func Slice[T any](slice []T, start, end int) []T {
	if start < 0 {
		start = 0
	}
	if end > len(slice) {
		end = len(slice)
	}
	if start >= end {
		return nil
	}
	return slice[start:end]
}
