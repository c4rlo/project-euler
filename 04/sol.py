#!/usr/bin/python3

import sys

m = 0

for i in range(999, 99, -1):
    for j in range(999, i - 1, -1):
        n = i * j
        ns = str(n)
        if ns == ns[::-1]:
            m = max(m, n)

print("Answer:", m)
