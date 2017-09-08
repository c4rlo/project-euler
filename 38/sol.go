package main

import "fmt"

func main() {
	maxProdNum := 0

	for n := 1; n <= 9876; n++ {
		var storage [10]byte
		prod := storage[:0]
		var forbidDigit [10]bool
		forbidDigit[0] = true
		m := 0
		for i := 1; ; i++ {
			m += n
			prod = appendNewDigits(prod, m, &forbidDigit)
			prodLen := len(prod)
			if prod == nil || prodLen > 9 {
				break
			}
			if prodLen == 9 {
				prodNum := parseDigits(prod)
				fmt.Printf("%v @ %v = %v\n", n, i, prodNum)
				if prodNum > maxProdNum {
					maxProdNum = prodNum
				}
				break
			}
		}
	}

	fmt.Printf("\n*** Winner: %v\n", maxProdNum)
}

func appendNewDigits(s []byte, n int, forbidDigit *[10]bool) []byte {
	var storage [5]byte
	temp := storage[:0]
	for k := n; k > 0; k /= 10 {
		digit := byte(k % 10)
		if forbidDigit[digit] {
			return nil
		}
		temp = append(temp, digit)
		forbidDigit[digit] = true
	}
	for i := len(temp) - 1; i >= 0; i-- {
		s = append(s, temp[i])
	}
	return s
}

func parseDigits(s []byte) int {
	n := 0
	for _, d := range s {
		n *= 10
		n += int(d)
	}
	return n
}
