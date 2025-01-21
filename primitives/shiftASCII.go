package primitives

// данная функция берет элементы строки, и каждый смещает на шаг по кодам ASCII
func ShiftASCII(s string, step int) string {
	f := ""
	for i := 0; i < len(s); i++ {
		f += string(s[i] + byte(step))
	}
	return f
}
