package onelab_homework1_lib

func ChangeLetterCase(str string) string {
	arrRune := []rune(str)

	for i := range arrRune {
		if arrRune[i] >= 'A' && arrRune[i] <= 'Z' {
			arrRune[i] += 32
		} else if arrRune[i] >= 'a' && arrRune[i] <= 'z' {
			arrRune[i] -= 32
		}
	}

	return string(arrRune)
}
