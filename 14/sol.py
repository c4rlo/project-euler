#!/usr/bin/python3

def collatz(n):
    return n / 2 if n % 2 == 0 else 3*n + 1

cache = { 1: 1 }

def chainlen(n):
    if n not in cache:
        cache[n] = 1 + chainlen(collatz(n))
    return cache[n]

mlen = 0
mstart = 0

for i in range(1, 1000000):
    length = chainlen(i)
    if length > mlen:
        mlen = length
        mstart = i

print("chain", mstart, "length:", mlen)
