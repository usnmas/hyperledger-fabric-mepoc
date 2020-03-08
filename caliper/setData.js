/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

'use strict';

module.exports.info = 'Set initial device records';

//let txIndex = 0;
let sensorId = ['s001', 's002', 's003', 's004', 's005', 's006'];
let actuatorId = ['a001', 'a002'];
let bc, contx;

module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;

    return Promise.resolve();
};

module.exports.run = function() {
    //txIndex++;
    let sid = sensorId[Math.floor(Math.random() * sensorId.length)];
    let aid = actuatorId[Math.floor(Math.random() * actuatorId.length)];

    let args = {
        chaincodeFunction: 'setData',
        chaincodeArguments: [sid, aid]
    };

    return bc.invokeSmartContract(contx, 'mycc', 'v1.0', args, 10);
};

module.exports.end = function() {
    return Promise.resolve();
};
