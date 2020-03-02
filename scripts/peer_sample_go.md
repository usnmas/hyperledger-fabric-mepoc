**chaincode install**


root@c0bddbf3a304:/opt/gopath/src/github.com/newcc# peer chaincode install -p github.com/newcc -n mycc -v 1.0


**chaincode instantiate**


root@c0bddbf3a304:/opt/gopath/src/github.com/newcc# peer chaincode instantiate -n mycc -v 1.0 -c '{"Args":[]}' -C mychannel


**Insert Device Record to State DB**


root@c0bddbf3a304:/opt/gopath/src/github.com/newcc# peer chaincode invoke -n mycc -c '{"Args":["setData", "s001", "a001"]}' -C mychannel


**Query Device Data**


root@c0bddbf3a304:/opt/gopath/src/github.com/newcc# peer chaincode query -n mycc -c '{"Args":["queryData","s001"]}' -C mychannel

