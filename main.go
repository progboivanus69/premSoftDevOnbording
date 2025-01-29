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
	"reflect"
	"time"
)

// var ee1 = errors.New("error check")
var testS = "privet"
var ffff = 10 // переменная, которую используем с указателем
func main() {

	var brother = 10 // создали переменную
	// var brother2, brother3 int = 23, 45
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
	tester2 := structures.User{Age: 12, Gender: "male"}
	//tester3 := &tester2
	//tester4 := []int{1, 3, 5}
	//var tester5 structures.User
	var tester6 [5]int = [5]int{1, 2, 3, 4, 5}
	tester6Sl := tester6[1:]
	fmt.Println(cap(tester6Sl))

	tester7Sl := []int{1, 5, 6}
	tester8sl := tester7Sl[1:]
	tester7Sl = tester7Sl[1:]
	fmt.Println(tester7Sl, tester8sl)
	tester7Sl = tester7Sl[0:]
	fmt.Println(cap(tester7Sl))
	tester7Sl = append(tester7Sl, 1, 2, 5) // очень важно - если мы аппендим - мы создаем новый расширенный массив, в который перекладываем значения из старого
	fmt.Println(cap(tester7Sl))
	fmt.Println(tester7Sl, tester8sl)
	tester7Sl[0] = 2332342
	fmt.Println(tester7Sl, tester8sl[0]) // поэтому тут не видны изменения в tester8sl, так как он ссылается все еще на старый массив
	//fmt.Println(tester2.Age)
	tester2.Name = "fsdf"

	var tester10 []int
	fmt.Println(tester10 == nil)

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

	//fmt.Println(TestTest(10))
	//fmt.Println(TestTest(10), testS)

	// тестируем функции с указателями и без

	fmt.Println(updater(ffff)) // при вызове просто достаем значение переменной, и присваиваем его копию уже локальной переменной функции
	fmt.Println(ffff)          // видим что изменения не коснулись переменной ffff

	fmt.Println(updaterWithAddr(&ffff)) //
	fmt.Println(ffff)
	fmt.Println(updater(ffff))
	fmt.Println(reflect.TypeOf(ffff))
	fmt.Println(typeChecker(ffff))
	var oldSlice []int = []int{1, 3, 4}
	fmt.Println(oldSlice, &oldSlice[1])
	newSl := updaterForEveryth(oldSlice)
	fmt.Println(newSl, &newSl[1], &newSl[1] == &oldSlice[1]) // видимо что адреса совпадают потому что при копировании слайса из одного в другой копируется адрес на массив
	fmt.Println(products(1, 4, 6))

	newSll := []int{1, 4, 5}
	for iii := 0; iii < len(newSll); iii++ {

		println(iii, newSll[iii])
	}
	varT := []int{10, 4}
	fmt.Println(varT)
	fmt.Println(products(varT...)) // этот параметр раскладывает слайс/массив на отдельные элементы при добавлении в тот самый слайс параметров factors ...int, то есть он каждый элемент слайса извлекает из переменной varT, и кладет в другой слайс, работает только со срезами!
	var newM1 map[int]string = map[int]string{1: "test", 2: "dfgf", 3: "dfgf", 4: "dfgf"}
	for k, v := range newM1 {
		println(k, v)
	}

	vvvv := stri{10, "sdf"}
	fmt.Println(vvvv)

	vvvv.wow()

	newM1[234] = "TEST"
	fmt.Println(newM1[234])
	_, checkerr := newM1[2345]
	fmt.Println(checkerr)
	//var newM = map[int]string{}
	ggg := 10000000
	fmt.Println(ggg, &ggg)
	println(newM1 == nil, newM1)
	fff := clos(&ggg)
	fmt.Println(fff())
	fmt.Println(ggg)

	stels := bicycle{
		Chain: 10,
		transport: transport{
			Speed: 155,
		},
	}
	fmt.Println(stels.Speed)
	newSlSS := []int{1, 4, 6}
	fmt.Println(&newSlSS[0])

	vvvvvv := Vertexx{XX: 1.0, YY: 3.0}
	fmt.Println(vvvvvv)

	firstSlice := []int{1, 2, 3}
	//secondSlice := []int {4,5,6}
	fmt.Println(multiplier(firstSlice...)) // тут используется оператор развертывания, он используется только для среза
	// fffff(3, 5)
	var fffff func(int, int) int
	fffff = sum
	fmt.Println(fffff(3, 5))
	fffff = product // такая функция не будет компилироваться раньше с глобальными функциями, поэтому использовать ее до этого участка кода и вызывать нельзя
	fmt.Println(fffff(3, 5))
	fmt.Println(reflect.TypeOf(fffff)) // т
	iter := 0
	for iter < 10 {

		iter++
	}
}

type creature struct {
	name string
	age  int
}
type human struct {
	creature
	gender bool
}

type animal struct {
	creature
	legsCount int
}

type speakable interface {
	Speak() string
}

func (a *animal) Speak() string {
	if a.legsCount == 4 {
		return "Bark"
	}
	return "Shhh"
}

func (h *human) Speak() string {
	return "Hi"
}

type Vertexx struct {
	XX, YY float64
}

type transport struct {
	Wheels bool
	Speed  int
}

type bicycle struct {
	Chain int
	transport
}
type stri struct {
	Age  int
	Name string
}

func (x stri) wow() {
	fmt.Println(x.Age)
}

func clos(x *int) func() int {
	return func() int {
		*x++
		return *x
	}

}

func sum(x, y int) int {

	return x + y
}
func product(x, y int) int {
	return x * y
}

func updater(ffff int) int { //имя параметра совпадает с именем переменной выше, но это разные переменные
	ffff = 150 //
	return ffff
}

func typeChecker(ffff interface{}) interface{} { //тип интерфейс принимает любой тип данных, функция возвращает тип передаваемого значения
	ffff = reflect.TypeOf(ffff).Name()

	switch ffff.(type) {
	case string:
		fmt.Println("string")

	case int:
		fmt.Println("int")

	default:
		fmt.Println("unknown")
	}

	return ffff
}
func updaterForEveryth(fff []int) []int {
	newW := fff
	return newW
}
func updaterWithAddr(ffff *int) int { //тут же мы указываем при объявлении переменной тип - ссылка
	*ffff = 1500
	return *ffff
}

func products(factors ...int) int { // параметр factors ...int это параметр среза переменного числа аргументов, в данном случае это срез из произвольного кодичества элементов типа int. Если я укажу другой тип, то будут другие элементы в срез собираться
	p := 1
	for _, n := range factors {
		p *= n
	}
	return p
}

func multiplier(n ...int) int {
	multRes := 1
	for _, element := range n {
		multRes *= element
	}
	return multRes
}

//func TestTest(fff int) int {
//	testS = "2"
//	return fff
//}
