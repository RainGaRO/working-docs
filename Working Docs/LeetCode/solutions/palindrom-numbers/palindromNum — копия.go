package palindromnumbers

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	revStr := reverseString(str)

	return str == revStr
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	testNumbers := []int{121, -121, 10, 12321}

	for _, num := range testNumbers {
		fmt.Printf("Is %d a palindrome? %t\n", num, isPalindrome(num))
	}
}

//==================== MORE VERSIONS ( 0ms )================================================================

/*
func isPalindrome(x int) bool {
	// Отрицательные числа и числа, оканчивающиеся на 0 (кроме 0 самого) не могут быть палиндромами
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// Создаем переменную для хранения второй половины числа в обратном порядке
	reversedSecondHalf := 0

	for x > reversedSecondHalf {
		reversedSecondHalf = reversedSecondHalf*10 + x%10
		x /= 10
	}

	// Число является палиндромом, если оно равно своей второй половине в обратном порядке
	// Или если оно равно второй половине без учета средней цифры (для нечетной длины)
	return x == reversedSecondHalf || x == reversedSecondHalf/10
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	var final int
	for num > 0 {
		mod := num % 10
		final = final*10 + mod
		num = num / 10
	}
	return final == x
}
*/
