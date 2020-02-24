'use strict';
const { Contract} = require('fabric-contract-api');

class testContract extends Contract {

   async queryMarks(ctx,deviceId) {
    let marksAsBytes = await ctx.stub.getState(deviceId);
    if (!marksAsBytes || marksAsBytes.toString().length <= 0) {
      throw new Error('Device with this ID does not exist: ');
       }
      let marks=JSON.parse(marksAsBytes.toString());
      return JSON.stringify(marks);
     }

   async queryAllMarks(ctx) {
     const starkKey = 'i000';
     const endKey = 'i999';
     const iterator = await ctx.stub.getStateByRange(startKey, endKey);
     const allResults = [];
     while(true) {
       const res = await iterator.next();
       if (res.value && res.value.value.toString()){
         const Key = res.value.key;
         let Record;
         try {
           Record = JSON.parse(res.value.value.toString('utf8'));
         } catch (err) {
           console.log(err);
           Record = res.value.value.toString('utf8');
         }
         allResults.push({Key, Record});
       }
       if (res.done) {
         console.log("End of data");
         await iterator.close();
         return JSON.stringify(allResults);
       }
     }
   }

   async addMarks(ctx,deviceId,time_c,temp_c,check_c) {
    let marks={
       item1:time_c,
       item2:temp_c,
       item3:check_c
       };
    await ctx.stub.putState(deviceId,Buffer.from(JSON.stringify(marks)));
    console.log('Device Info added To the ledger Succesfully..');
  }

   async deleteMarks(ctx,deviceId) {
    await ctx.stub.deleteState(deviceId);
    console.log('Device Info deleted from the ledger Succesfully..');
    }
}

module.exports=testContract;
