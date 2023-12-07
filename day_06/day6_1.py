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
for x in range(numgames):
    wgames = 0
    tdata_x = int(timedata[x])
    for y in range(tdata_x):
        if y > 0:
            ddata_y = int(distdata[x])
            tdist = y * (tdata_x-y)
            if tdist > ddata_y:
                wgames = wgames + 1

    waystowin.append(wgames) 

totalSum = 1
for w in waystowin:
    totalSum *= w

print(f"TotalSum: {totalSum}\n")
