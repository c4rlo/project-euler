#!/usr/bin/python3

with open('names.txt') as x: f = x.read()

names = [ n.strip('"') for n in f.split(',') ]
names.sort()

def score(name):
    return sum(ord(c) - ord('A') + 1 for c in name)

print(sum((i + 1) * score(n) for i, n in enumerate(names)))
