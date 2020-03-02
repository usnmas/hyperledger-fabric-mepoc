/*
 * This Node JS file is for test data creation for PoC 
 * OUTPUT : "SensorID","ActuatorID","TimeStamp (Millisecond)","Temperature","Hash Value"
 */

// timestamp
var cTime = new Date();     // 2020-02-17T09:53:33.658Z
var ts = cTime.getTime();   // 1581933213658 (Millisecond)

// Checksum value creation
var sha256 = require("js-sha256");

var sid = "s001";
var aid = "a001";
var tem = "35";
var dat = sid + aid + ts + tem;
var has = sha256(dat);
var tsS = String(ts);
var msgArr = [sid, aid, tsS, tem, has];
console.log(JSON.stringify(msgArr));
