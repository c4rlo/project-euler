#!/usr/bin/python3

def numdays(mo, yr):
    if mo == 2:
        return 28 if yr % 4 != 0 else 29
    return [31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31][mo - 1]

def incweek(yr, mo, day):
    nd = numdays(mo, yr)
    day += 7
    if day > nd:
        mo += 1
        day -= nd
        if mo > 12:
            yr += 1
            mo = 1
    return yr, mo, day

# 7 Jan 1900 was a Sunday
yr, mo, day = 1900, 1, 7

count = 0

while True:
    yr, mo, day = incweek(yr, mo, day)
    if yr > 2000: break
    if yr > 1900 and day == 1: count += 1

print(count)
