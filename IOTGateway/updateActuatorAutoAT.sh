#!/bin/bash

for ((i=0; i<10; i++))
do
  ./updateActuatorAuto.sh >> update100_`date +%Y%m%d%H%M`.txt
  sleep 120
done
