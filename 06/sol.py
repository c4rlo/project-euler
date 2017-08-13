#!/usr/bin/python3

sumsq = sum(i*i for i in range(1, 101))
sqsum = (101*50)**2

print(sumsq, sqsum, sqsum - sumsq)
