package slices

// функция копирует слайс в другой
func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return []int{}
	} else if maxLen > len(src) {
		return src
	}
	copSrc := make([]int, maxLen)
	copy(copSrc, src)

	return copSrc
}
