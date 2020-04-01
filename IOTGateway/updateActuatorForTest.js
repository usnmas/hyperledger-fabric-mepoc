"use strict";

const path = require("path");
const sha256 = require("js-sha256");

// for test data arbitrary selection
const math = require("math");
const tempList = ["15", "25", "35"];

const { FileSystemWallet, Gateway } = require("fabric-network");
const ccpPath = path.resolve(__dirname, ".", "connection.json");

async function submain() {
  try {
    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet");
    const wallet = new FileSystemWallet(walletPath);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.exists("admin");
    if (!userExists) {
      console.log('ERROR: cannot find "admin" user in the wallet');
      return;
    }

    // Create a new gateway for connecting to our peer node.
    // asLocalhost FALSE to connect remotely
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

    // Test data arbitrary creation and argument generation for invoke tx
    let argIdx = Math.floor(Math.random() * 100);
    let sid = "s" + argIdx;
    let aid = "a" + Math.trunc(argIdx / 10);

    let ct = new Date();
    let ts = ct.getTime();
    let tsS = String(ts);

    let tem = tempList[Math.floor(Math.random() * tempList.length)];

    let dat = sid + aid + ts + tem;
    let has = sha256(dat);

    await contract.submitTransaction("updateData", sid, aid, tsS, tem, has);
    //console.log('Transaction has been successfully invoked');

    // Disconnect from the gateway.
    await gateway.disconnect();
  } catch (error) {
    console.error(`Failed to submit transaction: ${error}`);
    process.exit(1);
  }
}

function main() {
  // Check start time
  var ct0 = new Date();
  var t0 = ct0.getTime();

  submain();

  // Check end time and show the gap
  var ct1 = new Date();
  var t1 = ct1.getTime();
  console.log("Latency in milliseconds : " + (t1 - t0));
}

main();
