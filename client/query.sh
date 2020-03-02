#!/bin/bash

CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
source query.data

echo "Query State DB record by Device ID"
docker exec \
  -e CORE_PEER_LOCALMSPID=Org1MSP \
  -e CORE_PEER_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp \
  cli \
  peer chaincode query \
    -n mycc \
    -c '{"Args":["'$func'","'$id'"]}' \
    -C mychannel

