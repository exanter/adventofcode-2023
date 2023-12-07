#!/usr/bin/env python3

import sys

data = open(sys.argv[1]).read().strip()
lines = data.split("\n")

timedata = lines[0].split()
distdata = lines[1].split()

waystowin = []
blah = timedata.pop(0)
blah = distdata.pop(0)

numgames = len(timedata)
ctdata_s = ""
cddata_s = ""
wgames = 0
for x in range(numgames):
    ctdata_s += "%s" % (timedata[x])
    cddata_s += "%s" % (distdata[x])

ctdata = int(ctdata_s)
cddata = int(cddata_s)
for y in range(ctdata):
    if y > 0:
        tdist = y * (ctdata-y)
        if tdist > cddata:
            wgames = wgames + 1

print(f"Total Ways: {wgames}\n")
