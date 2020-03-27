// Query Data Entry : Both sensorID & actuatorID allowed as an argument

"use strict";

const path = require("path");

const { FileSystemWallet, Gateway } = require("fabric-network");
const ccpPath = path.resolve(__dirname, ".", "connection.json");

const arg = process.argv[2];

async function main() {
  try {
    //////////////////////////////////////////////////////
    // Check start time
    var ct0 = new Date();
    var t0 = ct0.getTime();
    //////////////////////////////////////////////////////

    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet");
    const wallet = new FileSystemWallet(walletPath);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.exists("admin");
    if (!userExists) {
      console.log('An identity for "admin" does not exist in the wallet');
      return;
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccpPath, {
      wallet,
      identity: "admin",
      discovery: { enabled: true, asLocalhost: false }
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork("mychannel");

    // Get the contract from the network.
    const contract = network.getContract("mycc");

    // Evaluate the specified transaction by the input argument
    const result = await contract.evaluateTransaction("queryData", arg);

    //////////////////////////////////////////////////////
    // Check end time and show the gap
    var ct1 = new Date();
    var t1 = ct1.getTime();
    console.log("Latency in milliseconds : " + (t1 - t0));
    //////////////////////////////////////////////////////

    console.log(
      `Transaction has been evaluated, result is: ${result.toString()}`
    );
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    process.exit(1);
  }
}

main();
