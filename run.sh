#!/bin/bash

echo Python 2
time python2 -c 'from itertools import imap;print(sum(imap(int,open("BIGDATA2"))))'

echo AWK
time awk '{i+=$1} END{print i}' BIGDATA2

echo Python 3
time python3 -c 'print(sum(map(int,open("BIGDATA2"))))'

echo Python 3 No Unicode
time python3 -c 'print(sum(map(int,open("BIGDATA2","rb"))))'

echo PHP
time php -d 'memory_limit = 2G' -r 'echo  array_sum(file("BIGDATA2"))."\n";echo number_format(memory_get_peak_usage(1)/1024/1024)."M";'

echo Go 
time go run sum.go

echo C
time ./sumc


