# Usage 


- The initial files, "channel.tx / genesis.block / Org1MSPanchors.tx" should be copied or created first in a subfolder "config". 


- "generate.sh" will create a subfolder "crypto-config" and relevant files there. 


- ".env" is necessary to create docker bridge network with a correct name. Otherwise, the bridge network name will be mingled and may cause error during chaincode instantiation. 


- The execution permission of shell scripts should be adjusted on OS. 
