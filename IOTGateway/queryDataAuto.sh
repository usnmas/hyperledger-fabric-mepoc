#!/bin/bash

#for ((i=0; i<10; i++))
#do
#    node queryData-wa.js a"$i" &
#    sleep 1
#done

for ((i=0; i<250; i++))
do
    node queryDataForTest.js &
done
