# Usage 


- "generate.sh" will create a subfolder "config", "crypto-config" and relevant files there. 


- ".env" is necessary to create docker bridge network with a correct name. Otherwise, the bridge network name will be mingled and may cause error during chaincode instantiation. 


- The execution permission of shell scripts should be adjusted on OS. 
