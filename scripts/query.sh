#!/bin/bash

sudo docker exec -e “CORE_PEER_LOCALMSPID=Org1MSP” \
-e “CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp” \
cli \
peer chaincode query -n mycc -c '{"Args":["queryData","s001"]}' -C mychannel

