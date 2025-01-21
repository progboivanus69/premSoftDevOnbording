package slices

import "fmt"

// ModifySlice функция для вставки элементов в слайс
func ModifySlice(nums []int) []int {
	nums[2] = 10           // элемент будет и в исходном слайсе
	nums = append(nums, 6) // элемент не добавится в исходный слайс
	fmt.Println(nums)
	return nums
}
