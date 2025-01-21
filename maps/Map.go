package maps

// Map Данная функция изменяет элементы слайса, добавляя к ним текст
func Map(strst []string, mapFunc func(s string) string) []string {
	newSl := make([]string, len(strst))
	for i, elem := range strst {
		println(strst, "фулл массив,", newSl[i], "индекс массива newSl")
		println(elem, "элемент массива strst до изменений")
		newSl[i] = mapFunc(elem)
		println(newSl[i], "элемент массива newSl после изменений")

	}
	return newSl
}
