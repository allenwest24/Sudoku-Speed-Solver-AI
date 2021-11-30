
import sys

old = sys.argv[1]

new = ''
odd = True
tally = 0
for ii in range(len(old)):
    if old[ii] == '[':
        new += '{'
    elif old[ii] == ']':
        new += '}'
    elif old[ii] == ' ':
        new += ', '
    else:
        new += old[ii]
print(new)

