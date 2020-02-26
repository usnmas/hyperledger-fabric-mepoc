# fabric-mepoc

Master of Engineering Proof-of-Concept Project


## Usage


If needed, clone the [hyperledger/fabric-samples](https://github.com/hyperledger/fabric-samples) repository. This repository includes "bin" and "basic-network" folder. 


**1. Install the Hyperledger Fabric platform-specific binaries for the version specified into the $POCHOME/bin**
- configtxgen
- configtxlator
- cryptogen
- discover
- idemixgen
- orderer
- peer
- fabric-ca-client


**2. Execute $POCHOME/basic-network/generate.sh**


This will generate your cryptographic files (certificate and key) under the folder $POCHOME/basic-network/config and $POCHOME/basic-network/crypto-config


**3. Modify path in $POCHOME/basic-network/connection.json based on the above certificate**


**4. Modify CA_KEYFILE in $POCHOME/basic-network/docker-compose.yml based on the above certificate**


## Reference 
> Reference for chaincode development : [ChaincodeStubInterface](https://godoc.org/github.com/hyperledger/fabric-chaincode-go/shim#ChaincodeStubInterface)
