#!/bin/bash

CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
source update.data

echo "Update State DB record by Device ID"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org1MSP \
  -e CORE_PEER_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp \
  cli \
  peer chaincode invoke \
    -n mycc \
    -c '{"Args":["'$func'","'$a1'","'$a2'","'$a3'","'$a4'","'$a5'"]}' \
    -C mychannel

