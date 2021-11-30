
import sys

old = sys.argv[1]

new = ''
for ii in range(len(old)):
    if old[ii] == '[':
        new += '{'
    elif old[ii] == ']':
        new += '}'
    elif old[ii] == '0' and (old[ii+1] == ',' or old[ii+1] == ']') and (old[ii-1] == ',' or old[ii-1] == '['):
        new += "byte('.')"
    else:
        new += old[ii]

print(new)
