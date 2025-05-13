package palindromstrings

import "fmt"

func IsPalindrome(input string) bool { //брабатывает строки в кодировке Unicode
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func main() {
	testCases := []string{"", "G", "bob", "otto", "汉字汉", "test"}

	for _, input := range testCases {
		result := IsPalindrome(input)
		fmt.Printf("%q is a palindrome: %t\n", input, result)
	}
}

//================================================================
/*
func isPalindrome(s string) bool {    //работаем с байтами - это предпочтительнее для работы с ASCII или однобайтными кодировками
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("sasha"))
}
*/
