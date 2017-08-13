#!/usr/bin/python3

import sys

# a^2 + b^2 = c^2
# a + b + c = 1000
# a < b < c
# a b c = ?

# a^2 + b^2 = (1000 - a - b)^2
#           = 1000000 - 2000 a - 2000 b + 2 a b + a^2 + b^2

# 2000 a + 2000 b = 2 a b + 1000000

# a b c = ?

# 1 < a < b < c < 998

for a in range(1, 999):
    for b in range(a, 999):
        if 2000*a + 2000*b == 2*a*b + 1000000:
            print("a =", a, "; b =", b, "; c =", 1000 - a - b, "; abc =",
                    a*b*(1000-a-b))
            sys.exit()
