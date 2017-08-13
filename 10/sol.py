#!/usr/bin/python3

from math import sqrt

N = 2000000
NSQ = int(sqrt(N))

notPrime = bytearray(N)

for i in range(2, NSQ):
    for j in range(2, N):
        ij = i * j
        if ij >= N: break
        notPrime[ij] = 1

notPrime[0] = notPrime[1] = 1

total = sum(map(lambda inp: inp[0] if not inp[1] else 0, enumerate(notPrime)))

print(total)
