import json
import sys

ob = json.load(sys.stdin)
nset = set(sys.argv[1])
out = []
for w in ob:
    if len(nset.intersection(w)) == 0:
        out.append(w)
json.dump(out, sys.stdout)
sys.stderr.write('{} in {} out\n'.format(len(ob), len(out)))
