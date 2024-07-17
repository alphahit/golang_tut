package main

import (
	"fmt"
	"math"
)

// Function to find the number of digits in a number
func findDigits(n int, count int) int {
	if n == 0 {
		return count
	}
	return findDigits(n/10, count+1)
}

// Recursive function to check if a number is a palindrome
func checkPalindrome(num int, digitCount int) bool {
	if num == 0 || digitCount <= 1 {
		return true
	}
	divisor := int(math.Pow(10, float64(digitCount-1)))
	firstDigit := num / divisor
	lastDigit := num % 10
	if firstDigit != lastDigit {
		return false
	}
	num = (num % divisor) / 10
	digitCount -= 2

	return checkPalindrome(num, digitCount)
}

// Function to check if a number is a palindrome
func isPalindrome(number int) bool {
	if number < 0 {
		return false
	}
	digitCount := findDigits(number, 0)
	return checkPalindrome(number, digitCount)
}

// Function to check if all numbers in an array are palindromes
func checkArrayPalindrome(arr []int) int {
	flag := 1
	for _, num := range arr {
		if !isPalindrome(num) {
			flag = 0
			break
		}
	}
	return flag
}

func main() {
	arr := []int{111, 222, 333, 444, 555}
	arr2 := []int{111, 121, 131, 20}

	fmt.Println(checkArrayPalindrome(arr))  // Output: 1
	fmt.Println(checkArrayPalindrome(arr2)) // Output: 0
}
