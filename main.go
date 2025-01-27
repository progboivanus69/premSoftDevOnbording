package main

import (
	"awesomeProject/maps"
	"awesomeProject/primitives"
	"awesomeProject/slices"
	"awesomeProject/structures"
	"awesomeProject/vars"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"time"
)

func main() {
	var brother = 10          // создали переменную
	brother, sister := 14, 34 // а здесь с помощью синтаксического сахара краткой формы объявления переменной мы старой переменной смогли переприсвоить значение а не создать ее, потому что слева еще была указана новая переменная
	//brother:= 14 // а тут будет ошибка что нет новой переменной слева
	fmt.Printf("The type of the following var is: %T,the value is: %v\n", vars.BasStr, vars.BasStr) // используем спецификаторы для вывода типа и значения переменной
	fmt.Printf("The type of the following var is: %T,the value is: %v\n", vars.R, vars.R)           // выводим преобразованную строку в срез рун
	fmt.Printf("The type of the following var is: %T,the value is: %v\n", vars.BB, vars.BB)         // выводим преобразованную строку в срез байт
	slices.SliceIterator(vars.BB)                                                                   // используем функцию для перебора байтов в строке
	fmt.Println(brother, sister)                                                                    //
	fmt.Println(slices.Mapa["иван"])                                                                // вывод значения мапы по ключу из пакета maps
	fmt.Println(slices.Mapa)                                                                        // полный вывод мапы по переменной
	// fmt.Println(maps.Mapa["ива"])       //./main.go:28:14: undefined: mapa - необъъявленная мапа
	next := make(map[int]string, 10) // создание пустой мапы через встроенную функцию make с заданной длиной ключей
	fmt.Println(len(next))           // выводим длину мапы - равна нулю, потому что мапа пустая
	fmt.Println(next)                // выводим мапу
	next[10] = "bruh"                // задает ключ 10: "bruh"
	fmt.Println(next[10])            // смотрим на значение по ключу - получаем bruh
	delete(next, 10)                 // используем функцию delete для удаления в указанной мапе значения по ключу
	fmt.Println(next[10])            // данного ключа со значением более нет
	fmt.Println(len(next))           // проверяем что длина после удаления равна 0

	tester := structures.User{
		Name:   "Ваc",
		Age:    18,
		Gender: "male",
	} // помещаем в переменную структуру типа User с наполненными данными, проходящими валидацию
	adm := structures.Admin{Inh: structures.User{
		Name:   "1",
		Age:    0,
		Gender: "",
	},
		IsSuperUser: true} //  // помещаем в переменную структуру типа Admin с наполненными данными, не проходящими валидацию
	bs, _ := json.Marshal(tester) // приводим к формату JSON согласно меткам структуру tester
	fmt.Println(string(bs))       // выводим приведенный к строке JSON
	bsb, _ := json.Marshal(adm)   // приводим к формату JSON согласно меткам структуру adm
	fmt.Println(string(bsb))      // выводим приведенный к строке JSON
	v := validator.New()          // задаем из пакета валидатор в переменную v
	fmt.Println(v.Struct(tester)) // запускаем валидацию первого JSON и видим успех (результат - пустая структура)
	fmt.Println(v.Struct(adm))    // запускаем валидацию первого JSON и видим ошибку валидации на трех полях

	fmt.Println(fmt.Sprintf("переменные из пакета: A - %d, B - %d, C - %d, D - %d, F - %d, G - %d, H - %d", vars.A, vars.B, vars.C, vars.D, vars.F, vars.G, vars.H)) // выв
	slices.ModifySlice(slices.Nu)                                                                                                                                    // вызываем функцию из пакета slices, а так же передаем из пакета переменных слайс, но не сохраняем измененный слайс в переменную

	fmt.Println(slices.Nu)                   // [1 2 10 4 5] видим что исходный слайс остался без изменений
	fmt.Println(slices.Remove(slices.Nu, 2)) // выполняем функцию из пакета слайсов для удаления элемента по индексу

	maps.Map(
		[]string{"bro", "yoo"},
		slices.MapFunc,
	) // вызываем функцию для преобразования элементов слайса

	fmt.Println(slices.IntsCopy([]int{4, 6, 2}, 3))
	fmt.Println(primitives.ShiftASCII("bro12", 1))
	fmt.Println(primitives.LatinLetters("sefwefsfef 2134535sdffsfefsef sefefEEFFSF"))
	fmt.Println(primitives.GenerateSelfStory("Вася", 39, 234.3))

	var nilForArr [2]int
	//if nilForArr == nil {
	//	fmt.Println("nil!")
	//}
	fmt.Println(nilForArr) // заполнится нулевым значением - нули в массиве, у массивов нет nil
	var nilForSlice []int  // заполнится нулевым значением - nil, если явно не использовать функцию типа make
	if nilForSlice == nil {
		fmt.Println("nil!")
	}
	fmt.Println(nilForSlice)
	var nilForMap map[int]string // заполнится нулевым значением - nil, если явно не использовать функцию типа make, по факту просто объявили но не инициализировали
	if nilForMap == nil {
		fmt.Println("nil!")
	}
	fmt.Println(nilForMap)
	var makeForSlice []int = make([]int, 0) // тут же мы присвоили пустой слайс
	if makeForSlice == nil {
		fmt.Println("nil!")
	}
	fmt.Println(makeForSlice)
	var makeForMap map[int]string = make(map[int]string) // тут же мы присвоили пустую мапу
	if makeForMap == nil {
		fmt.Println("nil!")
	}
	fmt.Println(makeForMap)

	_, ifKeyExists := makeForMap[123]      // мапа возвращает значение а так же флаг существования ключа
	fmt.Println(ifKeyExists)               // видим что такого ключа нет
	makeForMap[123] = "Bruh"               //добавили ключ
	_, ifKeyExistsAgain := makeForMap[123] // теперь проверяем повторно существование данного ключа
	fmt.Println(ifKeyExistsAgain)

	var x = 10
	var pp = &x // reference - получение адреса объекта x
	// это эквивалент записи var pp *int = &x, или pp := &x
	var yy = *pp // эта же переменная не является ссылкой, она запишет то, что получит через указатель, но сама она не является указателем. Один раз получили через указатель значение x, и сохранили
	fmt.Println(yy)
	*pp = 14
	fmt.Println(x, pp, yy) // 14 0x14000110e08 10 пример того, что поскольку pp указатель, то в нем лежит адрес переменной x
	// в гошке все косплексные типы, и переменные, которые не могут получить дфеолтное значение без инициализации - получают nil, и указатель - не исключение

	var f *int
	fmt.Println(f)             // увидим nil, потому что переменной типа указатель не инициализирован адрес (пытаемся поулучить содержимое указателя а не содержимое той переменной, на которую могли бы сослаться, поэтому разыменование * не используется)
	fmt.Println(f == nil)      // будет правдой
	fmt.Println(&f == nil, &f) // будет ложью, потому что у переменной f есть свой адрес хранения

	if _, err := os.Open("foo.ext"); err != nil {
		fmt.Println(err) // вывели ошибку, указывающую на проблему
	} else {
		fmt.Println("All is fine.")
	}
	hour := time.Now().Hour()
	fmt.Println(hour)
	switch {
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}

	switch { // если не класть параметр, то этот свитч работает как if, сравнение case идет со значением true, то есть если результат сравнения в case true, то это отработает
	case true:
		fmt.Println("I'm test")
	case false:
		fmt.Println("I'm writingППППП")
	}
}
