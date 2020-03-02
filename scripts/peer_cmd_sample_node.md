**chaincode install**


peer chaincode install -n mycc -v 1.0 -p "/opt/gopath/src/github.com/newcc" -l "node"


**chaincode instantiate (Node.js)**


peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n mycc -l "node" -v 1.0 -c '{"Args":[]}'


**Adding Marks of Device**


peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"addMarks","Args":["i001","12345678","25","11"]}'


**Querying Marks of Device**


peer chaincode query -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"queryMarks","Args":["i001"]}'


**Querying All Marks of Device**


peer chaincode query -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"queryAllMarks","Args":[]}'


root@003efb4a2419:/opt/gopath/src/github.com/hyperledger/fabric/peer# peer chaincode query -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"queryAllMarks"}'
Error: non-empty JSON chaincode parameters must contain the following keys: 'Args' or 'Function' and 'Args'


root@003efb4a2419:/opt/gopath/src/github.com/hyperledger/fabric/peer# peer chaincode query -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"queryAllMarks","Args":[]}'
Error: endorsement failure during query. response: status:500 message:"transaction returned with failure: ReferenceError: startKey is not defined"
root@003efb4a2419:/opt/gopath/src/github.com/hyperledger/fabric/peer#


**Deleting Marks of a device from Ledger**


peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"deleteMarks","Args":["i002"]}'


**Invoke Call from OS script file**


sudo docker exec -e “CORE_PEER_LOCALMSPID=Org1MSP” -e “CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp” cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"addMarks","Args":["i002","22222222","26","22"]}'


**Query Call from OS script file**


sudo docker exec -e “CORE_PEER_LOCALMSPID=Org1MSP” \
-e “CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp” \
cli \
peer chaincode query -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"queryMarks","Args":["i001"]}'


