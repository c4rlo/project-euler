#!/usr/bin/python3

DIGITS = [
        'one', 'two', 'three', 'four', 'five',
        'six', 'seven', 'eight', 'nine', 'ten',
        'eleven', 'twelve', 'thirteen', 'fourteen', 'fifteen',
        'sixteen', 'seventeen', 'eighteen', 'nineteen', 'twenty',
        ]

TENS = [
        'ten', 'twenty', 'thirty', 'forty', 'fifty',
        'sixty', 'seventy', 'eighty', 'ninety',
        ]

def spellnum(n):
    if n <= 20: return DIGITS[n - 1]
    if n < 100:
        s = TENS[n // 10 - 1]
        if n % 10 == 0: return s
        return s + "-" + DIGITS[n % 10 - 1]
    if n == 1000: return "one thousand"
    s = DIGITS[n // 100 - 1] + " hundred"
    if n % 100 == 0: return s
    return s + " and " + spellnum(n % 100)

count = 0

for i in range(1, 1001):
    count += len(spellnum(i).replace(' ', '').replace('-', ''))

print(count)
