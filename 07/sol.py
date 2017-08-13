#!/usr/bin/python3

import sys
from math import sqrt

N = 200000
NTH = 10001
NSQ = int(sqrt(N))

notPrime = bytearray(N)

nnp = 0

for i in range(2, NSQ):
    for j in range(2, N):
        ij = i * j
        if ij >= N: break
        notPrime[ij] = 1

notPrime[0] = notPrime[1] = 1

nth = 0
for i, np in enumerate(notPrime):
    if np == 0: nth += 1
    if nth == NTH:
        print(str(NTH) + "th prime number is", i)
        sys.exit()

print("Got up to", str(nth) + "th", "prime number")
