package main

import "fmt"

func main() {
	for denomTens := 1; denomTens <= 9; denomTens++ {
		for denomOnes := 1; denomOnes <= 9; denomOnes++ {
			denom := 10*denomTens + denomOnes
			var numerTens, numerOnes int
			// Case 1
			numerTens = denomTens
			for numerOnes = 1; numerOnes <= denomOnes; numerOnes++ {
				numer := 10*numerTens + numerOnes
				checkFractions(numer, denom, numerOnes, denomOnes)
			}
			// Case 2
			numerTens = denomOnes
			if numerTens <= denomTens {
				for numerOnes = 1; numerOnes <= 9; numerOnes++ {
					numer := 10*numerTens + numerOnes
					checkFractions(numer, denom, numerOnes, denomTens)
				}
			}
			// Case 3
			numerOnes = denomOnes
			for numerTens = 1; numerTens <= denomTens; numerTens++ {
				numer := 10*numerTens + numerOnes
				checkFractions(numer, denom, numerTens, denomTens)
			}
			// Case 4
			numerOnes = denomTens
			for numerTens = 1; numerTens <= denomTens; numerTens++ {
				numer := 10*numerTens + numerOnes
				checkFractions(numer, denom, numerTens, denomOnes)
			}
		}
	}
}

func checkFractions(numer1, denom1, numer2, denom2 int) {
	if numer1 >= denom1 {
		return
	}
	n1, d1 := cancelFrac(numer1, denom1)
	n2, d2 := cancelFrac(numer2, denom2)
	if n1 == n2 && d1 == d2 {
		fmt.Printf("%v/%v == %v/%v\n", numer1, denom1, numer2, denom2)
	}
}

// Precondition: numer < denom
func cancelFrac(numer, denom int) (n, d int) {
	for i := 2; i <= numer; i++ {
		for numer%i == 0 && denom%i == 0 {
			numer /= i
			denom /= i
		}
	}
	return numer, denom
}
