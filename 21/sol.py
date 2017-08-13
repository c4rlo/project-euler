#!/usr/bin/python3

MAX = 10000
MAX2 = 26000  # max(d(i), i = 1...MAX) = 25320

divisors = list({ 1 } for i in range(MAX2))

for i in range(2, MAX2):
    for j in range(2 * i, MAX2, i):
        divisors[j].add(i)

d = list(sum(div) for div in divisors)

print("Divisor precalc done")

isAmicable = list(False for i in range(MAX))

for i in range(2, MAX):
    if isAmicable[i]: continue
    di = d[i]
    if di != i and d[di] == i:
        isAmicable[i] = True
        if di < MAX: isAmicable[di] = True

print(sum(i for i, a in enumerate(isAmicable) if a))
