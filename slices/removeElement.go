package slices

import "fmt"

// Remove функция для удаления элемента в слайсе
func Remove(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}

	// nums[i] = nums[len(nums)-1]
	for n := i; n < len(nums)-1; n++ {
		nums[n] = nums[n+1]
		fmt.Println(n, nums[n])
	}
	return nums[:len(nums)-1]
}
