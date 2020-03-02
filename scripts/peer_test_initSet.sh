#!/bin/bash
stime=`date +%s%N`
echo $stime
for (( i = 0; i < 100; ++i ))
do
peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"Args":["setData", "s001", "a001"]}' 
done
ltime=`date +%s%N`
echo $ltime
expr $ltime - $stime
