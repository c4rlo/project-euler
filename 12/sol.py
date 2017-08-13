#!/usr/bin/python3

from operator import mul
from functools import reduce
from math import sqrt
from sys import exit

N = 500000
NSQ = int(sqrt(N))

NotPrime = bytearray(N)

for i in range(2, NSQ):
    for j in range(2, N):
        ij = i * j
        if ij >= N: break
        NotPrime[ij] = 1

NotPrime[0] = NotPrime[1] = 1

Primes = [ i for i, np in enumerate(NotPrime) if not np ]

def factor(n):
    nn = n
    if nn > N: raise Exception("number", nn, "too large")
    fac = { }
    for p in Primes:
        if p > nn: break
        while n % p == 0:
            fac[p] = fac.get(p, 0) + 1
            n //= p
    return fac

def numfactors(f):
    return reduce(mul, (c + 1 for c in f.values()), 1)

def rmfac(a, b):
    for k, v in b.items():
        a[k] -= v
        if a[k] == 0: del a[k]
        elif a[k] < 0: raise Exception("oops")

def addfac(a, b):
    for k, v in b.items():
        a[k] = a.get(k, 0) + v

print("precalc done")

tri = 3
currfac = { 3: 1 }

for i in range(3, 200000):
    # print("i:", i, "tri:", tri, "currfac:", currfac, "factor(i-1):", factor(i-1), "factor(i+1)", factor(i+1))
    addfac(currfac, factor(i+1))
    rmfac(currfac, factor(i-1))
    tri += i
    #print(currfac, factor(tri))
    #if currfac != factor(tri): print("awkward", currfac, factor(tri))
    numf = numfactors(currfac)
    if numf > 300:
        print(tri, "has", numf, "divisors:", currfac)
        if numf > 500: exit()
