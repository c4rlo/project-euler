#!/usr/bin/python3

import itertools
import sys

a, b = 1, 1

for r in itertools.count(1):
    if len(str(a)) == 1000:
        print("fib[{}] = {}".format(r, a))
        sys.exit()
    a, b = b, a + b
