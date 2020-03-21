'use strict';

const path = require('path');
const sha256 = require("js-sha256");

const { FileSystemWallet, Gateway } = require('fabric-network');
const ccpPath = path.resolve(__dirname, '.', 'connection.json');

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        
        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('admin');
        if (!userExists) {
            console.log('ERROR: cannot find "admin" user in the wallet');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        // asLocalhost FALSE to connect remotely 
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'admin', discovery: { enabled: true, asLocalhost: false } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('mycc');

        // Submit (Invoke) the specified transaction.
        // Update the actuator status with data value beyond threshold
        // Select a sensor
        const ct = new Date();
        const ts = ct.getTime();

        const sid = "s12";
        const aid = "a1";
        const tem = "35";
        const dat = sid + aid + ts + tem;
        const has = sha256(dat);
        const tsS = String(ts);
        //const msgArr = [sid, aid, tsS, tem, has];
        await contract.submitTransaction('updateData', sid, aid, tsS, tem, has); 

        console.log('Transaction has been successfully invoked');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
