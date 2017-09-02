package main

import (
	"fmt"
	"math"
)

func main() {
	truncatables := map[uint]bool{2: true, 3: true, 5: true, 7: true}
	pr := NewPrimes()
	num := 0
	total := uint(0)
	for i := 4; num < 11; i++ {
		p := pr.PrimeAt(i)
		if isTruncatable(p, &pr) {
			fmt.Println(p)
			truncatables[p] = true
			total += p
			num++
		}
	}
	fmt.Printf("\nTotal: %v\n", total)
}

func isTruncatable(n uint, pr *Primes) bool {
	for m := uint(10); ; m *= 10 {
		k := n % m
		if k == n {
			break
		}
		if !pr.IsPrime(k) {
			return false
		}
	}

	for n /= 10; n > 0; n /= 10 {
		if !pr.IsPrime(n) {
			return false
		}
	}

	return true
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

func (pr *Primes) PrimeAt(i int) uint {
	for i >= len(pr.primes) {
		pr.ensure(pr.max * 2)
	}
	return pr.primes[i]
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
