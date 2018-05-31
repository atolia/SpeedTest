#!/bin/bash

for i in `seq 1 100000`;do echo $RANDOM;done >> BIGDATA

for i in `seq 1 200`;do cat BIGDATA >> BIGDATA2;done

rm -f BIGDATA
