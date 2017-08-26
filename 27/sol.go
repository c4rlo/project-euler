package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

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

type Primes struct {
	primes      []uint
	isComposite []bool
	max         uint
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
	// fmt.Printf("Generated primes up to %v\n", max)
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

type Record struct {
	a    int
	b    int
	nMax int
}

func (rec *Record) Record(a, b, nMax int) {
	if nMax > rec.nMax {
		*rec = Record{a, b, nMax}
	}
}

func usage(argv0 string) {
	fmt.Fprintf(os.Stderr, "usage: %s A_MAX B_MAX\n", argv0)
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage(os.Args[0])
	}

	aMax_, err := strconv.ParseUint(os.Args[1], 10, strconv.IntSize)
	if err != nil {
		usage(os.Args[0])
	}
	bMax, err := strconv.ParseUint(os.Args[2], 10, strconv.IntSize)
	if err != nil {
		usage(os.Args[0])
	}
	aMin, aMax := -int(aMax_), int(aMax_)

	record := Record{}
	pr := NewPrimes()

	for _, b := range pr.PrimesUpToAtLeast(uint(bMax)) {
		if b > 1000 {
			break
		}
		// fmt.Printf("b = %v\n", b)
		b := int(b)
		for a := aMin; a <= aMax; a++ {
			// fmt.Printf("  a = %v\n", a)
			for n := 1; ; n++ {
				value := n*n + a*n + b
				if value > 0 && pr.IsPrime(uint(value)) {
					record.Record(a, b, n+1)
				} else {
					// fmt.Printf("    %v primes\n", n)
					break
				}
			}
		}
	}
	fmt.Printf("n^2 + %v*n + %v yields %v consecutive primes (a*b = %v)\n",
		record.a, record.b, record.nMax, record.a*record.b)
	fmt.Printf("Primes up to %v were calculated\n", pr.max)
}
