package util

// RemoveFirstNElements 去掉切片的前 n 个元素并返回剩余的切片
func RemoveFirstNElements[T any](s []T, n int) []T {
	if n > len(s) {
		return []T{} // 如果 n 大于切片长度，返回空切片
	}
	return s[n:]
}