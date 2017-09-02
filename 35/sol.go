package main

import (
	"fmt"
	"math"
)

const MILLION uint = 1000000

func main() {
	num := 0
	pr := NewPrimes()
	for _, p := range pr.PrimesUpToAtLeast(MILLION - 1) {
		if p >= MILLION {
			break
		}
		if isCircularPrime(p, &pr) {
			num++
		}
	}
	fmt.Printf("There are %v circular primes below %v.\n", num, MILLION)
}

func isCircularPrime(p uint, pr *Primes) bool {
	for n := rotate(p); n != p; n = rotate(n) {
		if !pr.IsPrime(n) {
			return false
		}
	}
	return true
}

func rotate(n uint) uint {
	d := n % 10
	n /= 10
	for k := n; k > 0; k /= 10 {
		d *= 10
	}
	return d + n
}

type Primes struct {
	primes      []uint
	isComposite []bool
	max         uint
}

func NewPrimes() Primes {
	return Primes{
		primes:      []uint{2, 3, 5},
		isComposite: []bool{true, true, false, false, true, false},
		max:         5,
	}
}

func (pr *Primes) PrimesUpToAtLeast(max uint) []uint {
	pr.ensure(max)
	return pr.primes
}

func (pr *Primes) IsPrime(n uint) bool {
	pr.ensure(n)
	return !pr.isComposite[n]
}

func (pr *Primes) ensure(max uint) {
	if max <= pr.max {
		return
	}
	pr.isComposite = ensureLen(pr.isComposite, int(max)+1)
	sqrtMax := uint(math.Sqrt(float64(max)) + 1)
	for _, p := range pr.PrimesUpToAtLeast(sqrtMax) {
		if p > sqrtMax {
			break
		}
		iMin := pr.max + p - (pr.max % p)
		for i := iMin; i <= max; i += p {
			pr.isComposite[i] = true
		}
	}
	for i := pr.max + 1; i <= max; i++ {
		if !pr.isComposite[i] {
			pr.primes = append(pr.primes, i)
		}
	}
	pr.max = max
}

func ensureLen(slice []bool, length int) []bool {
	if length > cap(slice) {
		tmp := make([]bool, length)
		copy(tmp, slice)
		slice = tmp
	}
	if length > len(slice) {
		return slice[:length]
	} else {
		return slice
	}
}
