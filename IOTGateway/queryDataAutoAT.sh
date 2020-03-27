#!/bin/bash

for ((i=0; i<10; i++))
do
  ./queryDataAuto.sh >> query250_`date +%Y%m%d%H%M`.log
  sleep 150
done
