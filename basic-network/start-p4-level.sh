#!/bin/bash
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose-p4.yaml down
docker-compose -f docker-compose-p4.yaml up -d
docker ps -a

# wait for Hyperledger Fabric to start
export FABRIC_START_TIMEOUT=10
sleep ${FABRIC_START_TIMEOUT}

ORG_MSP_VAR="CORE_PEER_LOCALMSPID=Org1MSP"
CLI_MSP_VAR="CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp"

PEER[0]="CORE_PEER_ADDRESS=peer0.org1.example.com:7051"
PEER[1]="CORE_PEER_ADDRESS=peer1.org1.example.com:8051"
PEER[2]="CORE_PEER_ADDRESS=peer2.org1.example.com:9051"
PEER[3]="CORE_PEER_ADDRESS=peer3.org1.example.com:10051"

# Create the channel
#docker exec -e ${ORG_MSP_VAR} -e ${PEER_MSP_VAR} peer0.org1.example.com peer channel create -o orderer.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/channel.tx
docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e ${PEER[0]} cli peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/channel.tx

# Channel block fetch
#sudo docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer1.org1.example.com peer channel fetch newest mychannel.block -c mychannel -o orderer.example.com:7050
# Join the channel
#sudo docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer1.org1.example.com peer channel join -b mychannel.block
#sudo docker exec -e ${ORG_MSP_VAR} -e ${PEER_MSP_VAR} peer0.org1.example.com peer channel join -b mychannel.block

# Join the channel
#docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e ${PEER0} cli peer channel join -b mychannel.block
for ((i=0; i<4; i++))
do
  docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e "${PEER[i]}" cli peer channel join -b mychannel.block
done

# Install the chaincode
#docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e ${PEER0} cli peer chaincode install -p github.com/chaincode/newcc -n mycc -v 1.0
for ((i=0; i<4; i++))
do
  docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e "${PEER[i]}" cli peer chaincode install -p github.com/chaincode/newcc -n mycc -v 1.0
done

# Instantiate the chaincode
docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e ${PEER[0]} cli peer chaincode instantiate -C mychannel -n mycc -v 1.0 -c '{"Args":[]}'
docker exec -e ${ORG_MSP_VAR} -e ${CLI_MSP_VAR} -e ${PEER[2]} cli peer chaincode instantiate -C mychannel -n mycc -v 1.0 -c '{"Args":[]}'
