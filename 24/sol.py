#!/usr/bin/python3

import sys

n = 0
N = 1000000

def perm(prefix, segment):
    global n
    if len(prefix) == 10:
        n += 1
        if n == N:
            print(prefix.decode())
            sys.exit()
    else:
        newp = bytearray(prefix)
        newp.append(0)
        last = len(prefix)
        for i, b in enumerate(segment):
            newp[last] = b
            perm(newp, segment[:i] + segment[i+1:])

perm(b"", b"0123456789")
