#!/usr/bin/python3

MAX = 28123

divisors = [ { 1 } for i in range(MAX) ]
divisors[0] = { }

for i in range(2, MAX):
    for j in range(2 * i, MAX, i):
        divisors[j].add(i)

print("Divisors precalc done")

d = [ sum(div) for div in divisors ]
abundant = [ i for i in range(MAX) if d[i] > i ]

print("Abundant precalc done")

cbw = [ False for i in range(MAX) ]
#absum = [ False for i in range(MAX) ]

for i, a in enumerate(abundant):
    for j in range(i + 1):
        s = a + abundant[j]
        if s >= MAX: break
        cbw[s] = True
        #absum[s] = (a, abundant[j])

#for i in range(MAX):
#    if d[i] < i: print('<', end='')
#    elif d[i] == i: print('=', end='')
#    else: print('>', end='')
#    print("", i, "", end='')
#    if cbw[i]: print(absum[i])
#    else: print()

print(sum(i for i, c in enumerate(cbw) if not c))
