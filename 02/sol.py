#!/usr/bin/python3

tot = 0

a, b = 0, 1

while True:
    if b >= 4000000: break
    if b % 2 == 0: tot += b 
    a, b = b, a + b

print("Answer is", tot)
