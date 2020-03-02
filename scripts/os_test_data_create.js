/*
 *  This Node JS file is for test data creation for PoC 
 *  This needs an argument for temperature data
 *  OUTPUT : "SensorID","ActuatorID","TimeStamp (Millisecond)","Temperature","Hash Value"
 *    
*/

// timestamp
var cTime = new Date();     // 2020-02-17T09:53:33.658Z
var ts = cTime.getTime();   // 1581933213658 (Millisecond)

// Checksum value creation
var sha256 = require("js-sha256");

// Adapt temperature as input parameter
process.argv.forEach((val, index) => {}); 

var sid = "s001";
var aid = "a001";
//var tem = "35";
var tem = process.argv[2]; 
var dat = sid + aid + ts + tem;
var has = sha256(dat);
var tsS = String(ts);
var msgArr = [sid, aid, tsS, tem, has];

console.log(JSON.stringify(msgArr));

console.log("a1="+sid); 
console.log("a2="+aid); 
console.log("a3="+tsS); 
console.log("a4="+tem); 
console.log("a5="+has); 

