/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const path = require('path');
const math = require('math');
const { FileSystemWallet, Gateway } = require('fabric-network');
const ccpPath = path.resolve(__dirname, '..', 'basic-network', 'connection.json');

async function main() {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('admin');
        if (!userExists) {
            console.log('ERROR: cannot find "admin" user in the wallet');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'admin', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('mycc');

        // Submit the specified transaction.
        // Create sensor ID & actuator ID 
        // Assign every 10 sensors to an actuator
        let i, j, sid, aid; 
        for (i=0; i<100; i++) {
            sid = "s" + i; 
            j = math.trunc(i/10); 
            aid = "a" + j; 
            await contract.submitTransaction('setData',sid,aid); 
        }

        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
