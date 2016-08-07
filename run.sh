#!/bin/bash

echo -n Python 2 result=
time python2 -c 'from itertools import imap;print(sum(imap(int,open("BIGDATA2"))))'

echo -----------------------
echo -n AWK result=
time awk '{i+=$1} END{print i}' BIGDATA2

echo -----------------------
echo -n Python 3 result=
time python3 -c 'print(sum(map(int,open("BIGDATA2"))))'

echo -----------------------
echo -n Python 3 No Unicode result=
time python3 -c 'print(sum(map(int,open("BIGDATA2","rb"))))'

echo -----------------------
echo -n PHP result=
time php -d 'memory_limit = 2G' -r 'echo  array_sum(file("BIGDATA2"))."\n";'

echo -----------------------
echo -n Go result=
time go run sum.go

echo -----------------------
echo -n C result=
time ./sumc


