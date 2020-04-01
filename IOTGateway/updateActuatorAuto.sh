#!/bin/bash

for ((i=0; i<100; i++))
do
    node updateActuatorForTest.js &
done
