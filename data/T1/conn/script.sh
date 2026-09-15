#!/bin/bash

Nsample=583

for ((i = 0; i < Nsample; i++))
do
    cat ${i}.dat | awk '175 <= NR && NR <= 238 { printf("%s\n", $0) }' > tmp.dat
    mv tmp.dat ${i}.dat
done
