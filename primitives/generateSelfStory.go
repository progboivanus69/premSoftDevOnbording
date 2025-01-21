package primitives

import "fmt"

// данная функция использует шаблон с плейсхолдерами
func GenerateSelfStory(name string, age int, money float64) string {
	form := fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
	return form
}
