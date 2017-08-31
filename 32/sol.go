package main

import "fmt"

var matchingProducts map[int]bool = make(map[int]bool)

func main() {
	findProducts(2, 3)
	findProducts(1, 4)
	sum := 0
	for key := range matchingProducts {
		sum += key
	}
	fmt.Printf("Found %v unique products with a sum of %v\n",
		len(matchingProducts), sum)
}

func findProducts(len1, len2 int) {
	digits := [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	permutations(digits[:], len1, func(prefix, segment []byte) {
		a := toNum(prefix)
		permutations(segment, len2, func(prefix2, segment2 []byte) {
			b := toNum(prefix2)
			c := a * b
			if hasAllDigits(c, segment2) {
				fmt.Printf("%v * %v = %v\n", a, b, c)
				matchingProducts[c] = true
			}
		})
	})
}

func permutations(segment []byte, targetLen int, f func(_, _ []byte)) {
	perms(nil, segment, targetLen, f)
}

func perms(prefix []byte, segment []byte, targetLen int, f func(_, _ []byte)) {
	if len(prefix) == targetLen {
		f(prefix, segment)
		return
	}
	newPrefix := make([]byte, len(prefix)+1)
	copy(newPrefix, prefix)
	last := len(prefix)
	newSegment := make([]byte, len(segment)-1)
	for i, v := range segment {
		newPrefix[last] = v
		copy(newSegment, segment[:i])
		copy(newSegment[i:], segment[i+1:])
		perms(newPrefix, newSegment, targetLen, f)
	}
}

func toNum(s []byte) int {
	result := 0
	for _, b := range s {
		result *= 10
		result += int(b)
	}
	return result
}

func hasAllDigits(n int, digits []byte) bool {
	nDigits := make([]byte, 0, 8)
	for n > 0 {
		nDigits = append(nDigits, byte(n%10))
		n /= 10
	}
	if len(nDigits) != len(digits) {
		return false
	}
	for _, d := range digits {
		found := false
		for _, nd := range nDigits {
			if d == nd {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
