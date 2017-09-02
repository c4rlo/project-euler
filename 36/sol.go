package main

import "fmt"

func main() {
	total := 0

	var digitsArray [6]byte

	// Enumerate even-length base-10 palindromes
	for i := 1; i < 1000; i++ {
		digits := appendDigits(digitsArray[:0], i)
		digits = appendReverseDigits(digits, i)
		total += check(digits)
	}

	// Enumerate odd-length base-10 palindromes
	for i := 0; i < 100; i++ {
		digits := appendDigits(digitsArray[:0], i)
		for d := byte(1); d <= 9; d++ {
			digits2 := append(digits, d)
			digits2 = appendReverseDigits(digits2, i)
			total += check(digits2)
		}
	}

	fmt.Printf("\nTotal: %v\n", total)
}

func check(digits []byte) int {
	n := fromDigits(digits)
	if isBinaryPalindrome(n) {
		fmt.Println(n)
		return n
	} else {
		return 0
	}
}

func appendDigits(digits []byte, n int) []byte {
	var storage [6]byte
	temp := storage[:0]
	for k := n; k > 0; k /= 10 {
		temp = append(temp, byte(k%10))
	}
	for i := len(temp) - 1; i >= 0; i-- {
		digits = append(digits, temp[i])
	}
	return digits
}

func appendReverseDigits(digits []byte, n int) []byte {
	for k := n; k > 0; k /= 10 {
		digits = append(digits, byte(k%10))
	}
	return digits
}

func fromDigits(digits []byte) int {
	n := 0
	for _, d := range digits {
		n = 10*n + int(d)
	}
	return n
}

func isBinaryPalindrome(n int) bool {
	var length uint
	for m := n; m > 0; m >>= 1 {
		length++
	}
	for i := uint(0); i < length/2; i++ {
		if (n>>i)&1 != (n>>(length-1-i))&1 {
			return false
		}
	}
	return true
}
