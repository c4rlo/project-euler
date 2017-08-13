#!/usr/bin/python3

from math import sqrt

NUM = 600851475143

#NUM = 10005

NUMSQ = int(sqrt(NUM))
NUMSQSQ = int(sqrt(sqrt(NUM)))

notPrime = bytearray(NUMSQ)

for i in range(2, NUMSQSQ):
    for j in (i * k for k in range(2, NUMSQ)):
        if j >= NUMSQ: break
        notPrime[j] = True

#for i, np in enumerate(notPrime):
#    print(i, "is not prime" if np else "is prime")

#for i, np in reversed(enumerate(notPrime)):
for i, np in zip(range(NUMSQ - 1, -1, -1), reversed(notPrime)):
    #print(i, np, NUM % i)
    if not np and NUM % i == 0:
        print("Largest prime factor of", NUM, "is", i)
        break
