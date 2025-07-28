package _0250712

import (
	"fmt"
	"testing"
	"unicode"
)

// convert uppercase -> lowercase
// remove all non-alphanumeric: not both number and letter

func isPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		for left < right && isAlphaNumeric(runes[left]) {
			left++
		}
		for left < right && isAlphaNumeric(runes[right]) {
			right--
		}
		if toLower(runes[left]) != toLower(runes[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r)
}

func toLower(r rune) rune {
	return unicode.ToLower(r)
}

func Test_isPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

/*
Nếu bạn không muốn dùng unicode package trong Go
(ví dụ: muốn tự xử lý ký tự ASCII),
bạn hoàn toàn có thể viết lại hàm isPalindrome sử dụng ASCII thuần túy
— chỉ xử lý chữ cái (a–z, A–Z) và chữ số (0–9).
*/
