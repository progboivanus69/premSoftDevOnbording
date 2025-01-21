package slices

// Данная функция добавляет суффикс к изменяемому элементу мапы
func MapFunc(s string) string {
	s += "Changed"
	return s
}
