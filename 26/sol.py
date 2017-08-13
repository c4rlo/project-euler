#!/usr/bin/python3

def fracred(numer, denom):
    m = min(numer, denom)
    for i in range(2, m):
        if i*i > m:
            break
        if numer % i == 0 and denom % i == 0:
            numer //= i
            denom //= i
    return numer, denom

def closest_digit(numer, denom):
    for i in range(9, -1, -1):
        diff = numer*10 - i*denom
        if diff >= 0:
            return i, (diff == 0)
    assert "oops: {}/{}".format(numer, denom)

def as_cycdec(d):
    numer, denom = 1, d
    prefix, cont = "", ""
    frac2index = {}
    while True:
        numer, denom = fracred(numer, denom)
        index = frac2index.get((numer, denom))
        if index is not None:
            cont = prefix[index:]
            prefix = prefix[:index]
            if cont == "0":
                cont = ""
            break
        digit, is_exact = closest_digit(numer, denom)
        frac2index[(numer, denom)] = len(prefix)
        prefix += str(digit)
        if is_exact:
            break
        numer, denom = 10*numer - denom*digit, denom
    return prefix, cont

max_cycle_len = 0
max_cycle_num = None

for d in range(1, 1000):
    prefix, cont = as_cycdec(d)
    cycle_len = len(cont)
    if cycle_len > max_cycle_len:
        max_cycle_len = cycle_len
        max_cycle_num = (d, prefix, cont)
    if d <= 10:
        if len(cont) == 0:
            print("1/{} = 0.{}".format(d, prefix))
        else:
            print("1/{} = 0.{}({})".format(d, prefix, cont))

print("1/{} = 0.{}({}), cycle len = {}".format(*max_cycle_num, max_cycle_len))
