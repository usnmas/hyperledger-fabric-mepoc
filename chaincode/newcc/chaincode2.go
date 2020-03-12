/*
 * SPDX-License-Identifier: Apache-2.0
 * UNDER THE AUTHORITY OF CHUNG YUP KIM & ECS VUW
 */

package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	//sc "github.com/hyperledger/fabric-protos-go/peer"
)

// POCChaincode is the definition of the chaincode structure.
type POCChaincode struct {
}

// POCSensorData is the definition of the sensor data structure.
type POCSensorData struct {
	SensorID    string `json:"sid"`
	ActuatorID  string `json:"aid"`
	UnixTime    string `json:"utime"`
	Temperature string `json:"temp"`
	HashValue   string `json:"hashv"`
}

// POCDevice is the definition of device list.
type POCDevice struct {
	SensorID   string `json:"sid"`
	ActuatorID string `json:"aid"`
}

// POCActData is the definition of actuator status.
type POCActData struct {
	ActuatorID string `json:"aid"`
	Status     string `json:"status"`
}

// Prescribed Threshold
const MinThres, MaxThres = 20, 30

// Prescribed Actuator Status
const ActOn, ActOff = "ON", "OFF"

// Prescribed Delayed Time Allowed in seconds
const TimeThres = 60

// Init() is called when the chaincode is instantiated by the blockchain network.
func (cc *POCChaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	_, params := stub.GetFunctionAndParameters()
	fmt.Println("Init() is called with params: ", params)

	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
// Params : functionName, parameters //
func (cc *POCChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()

	switch fcn {
	case "setDevice":
		return cc.setDevice(stub, params)
	//case "setData":
	//	return cc.setData(stub, params)
	case "queryData":
		return cc.queryData(stub, params)
	case "queryAllData":
		return cc.queryAllData(stub, params)
	case "updateData":
		return cc.updateData(stub, params)
	default:
		return shim.Error("Input function name error")
	}
}

// INVOKE : Set state database with new device and actuator //
// Params : SensorID, ActuatorID //
// Return : Success //
func (cc *POCChaincode) setDevice(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// number of params check
	if len(params) != 2 {
		return shim.Error("Incorrect no. of parameters (setDevice)")
	}
	sensorID, actuatorID := params[0], params[1]

	err := stub.PutState(sensorID, []byte(actuatorID))
	if err != nil {
		return shim.Error("failed to putState for set Device : " + err.Error())
	}

	err = stub.PutState(actuatorID, []byte(ActOff))
	if err != nil {
		return shim.Error("failed to putState for set actuator : " + err.Error())
	}

	/*
		// make device db structure
		device := &POCDevice{SensorID: sensorID, ActuatorID: actuatorID}
		deviceBytes, err := json.Marshal(device)
		if err != nil {
			return shim.Error("failed to marshal for set Device : " + err.Error())
		}

		// save database with new device info.
		err = stub.PutState(sensorID, deviceBytes)
		if err != nil {
			return shim.Error("failed to putState for set Device : " + err.Error())
		}

		// make actuator db structure
		actuator := &POCActData{ActuatorID: actuatorID, Status: ActOff}
		actuatorBytes, err := json.Marshal(actuator)
		if err != nil {
			return shim.Error("failed to marshal for set actuator : " + err.Error())
		}

		// save database with new device info.
		err = stub.PutState(actuatorID, actuatorBytes)
		if err != nil {
			return shim.Error("failed to putState for set actuator : " + err.Error())
		}
	*/

	return shim.Success([]byte("New device record setting OK"))
}

/*
// Since setDevice created for db separation, setData is deprecated //
// INVOKE : Set state database with new device info. //
// Params : SensorID, ActuatorID //
// Return : Success //
func (cc *POCChaincode) setData(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// number of params check
	if len(params) != 2 {
		return shim.Error("Incorrect no. of parameters (setData)")
	}
	sensorID, actuatorID := params[0], params[1]

	// make database structure
	actuator := &POCActData{SensorID: sensorID, ActuatorID: actuatorID, Status: ActOff}
	actuatorBytes, err := json.Marshal(actuator)
	if err != nil {
		return shim.Error("failed to marshal for setData : " + err.Error())
	}

	// save database with new device info.
	err = stub.PutState(sensorID, actuatorBytes)
	if err != nil {
		return shim.Error("failed to putState for setData : " + err.Error())
	}

	return shim.Success([]byte("New record setting OK"))
}
*/

// QUERY : Query records //
// Params : sensorID or actuatorID //
// Return : mapped actuator ID or actuator status respectively //
func (cc *POCChaincode) queryData(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// number of params check
	if len(params) != 1 {
		return shim.Error("Incorrect no. of parameters (queryData)")
	}
	id := params[0]

	// query status
	queryStatus, err := stub.GetState(id)
	if err != nil {
		return shim.Error("failed to query status, error: " + err.Error())
	}
	if queryStatus == nil {
		return shim.Error("failed to query status which is null")
	}

	return shim.Success(queryStatus)
}

// QUERY : Query all the records in DB //
// Params : startkey & endkey (either sensorID or actuatorID) //
// Return : All records //
func (cc *POCChaincode) queryAllData(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	//startKey := "s0"
	//endKey := "s999"

	// number of params check
	if len(params) != 2 {
		return shim.Error("Incorrect no. of parameters (queryData)")
	}
	startKey, endKey := params[0], params[1]

	// query all the status
	recordIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error("failed to query all records, error: " + err.Error())
	}

	defer recordIterator.Close()

	if recordIterator.HasNext() {
		for recordIterator.HasNext() {
			queryResult, err := recordIterator.Next()
			if err != nil {
				return shim.Error("failed to fetch data, error: " + err.Error())
			}
			// can be checked in the cc container
			fmt.Println(string(queryResult.GetKey()), ":", string(queryResult.GetValue()))
		}
	}

	return shim.Success([]byte("Query all records successful"))
}

// INVOKE : Update actuator status //
// Params : SensorID, ActuatorID, UnixTime, Temperature, HashValue //
// Return : Success //
func (cc *POCChaincode) updateData(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// number of params check
	if len(params) != 5 {
		return shim.Error("Incorrect no. of parameters (updateData)")
	}
	sensorID, actuatorID, unixTime, temperature, hashValue := params[0], params[1], params[2], params[3], params[4]

	// check data integrity & latency
	sensorData := &POCSensorData{SensorID: sensorID, ActuatorID: actuatorID, UnixTime: unixTime, Temperature: temperature, HashValue: hashValue}
	sensorSrc := sensorID + actuatorID + unixTime + temperature
	sensorData.checkData(sensorSrc)

	// check whether temperature is integer
	tempInt, err := strconv.Atoi(temperature)
	if err != nil {
		return shim.Error("Input temperature is not integer")
	}

	// get db status
	/*
		actuatorBytes, err := stub.GetState(actuatorID)
		if err != nil {
			return shim.Error("failed to get queryData for updateData, error: " + err.Error())
		}
		actuator := POCActData{}
		err = json.Unmarshal(actuatorBytes, &actuator)
		if err != nil {
			return shim.Error("failed to unmarshal for updateData, error: " + err.Error())
		}
	*/
	actuatorStatus, err := stub.GetState(actuatorID)
	if err != nil {
		return shim.Error("failed to get queryData for updateData, error: " + err.Error())
	}
	actuatorStatusST := string(actuatorStatus)

	// check threshold & status
	switch {
	case tempInt > MaxThres && actuatorStatusST == ActOn:
		return shim.Success([]byte("Actuator already ON"))
	case tempInt > MaxThres && actuatorStatusST == ActOff:
		actuatorStatusST = ActOn
		err = stub.PutState(actuatorID, []byte(actuatorStatusST))
		if err != nil {
			return shim.Error("failed to update status to ON : " + err.Error())
		}
		err = stub.SetEvent("updateEvent", []byte(actuatorStatusST))
		if err != nil {
			return shim.Error("failed to emit event: " + err.Error())
		}
	case tempInt < MinThres && actuatorStatusST == ActOn:
		actuatorStatusST = ActOff
		err = stub.PutState(actuatorID, []byte(actuatorStatusST))
		if err != nil {
			return shim.Error("failed to update status to OFF : " + err.Error())
		}
		err = stub.SetEvent("updateEvent", []byte(actuatorStatusST))
		if err != nil {
			return shim.Error("failed to emit event: " + err.Error())
		}
	case tempInt < MinThres && actuatorStatusST == ActOff:
		return shim.Success([]byte("Actuator already OFF"))
	default:
		return shim.Success([]byte("Actuator does not need to be adjusted"))
	}

	return shim.Success([]byte("Actuator Status Set OK"))
}

// inner-call function module : check sensor data //
// Params : //
// Return : if error, code 11 & 12 //
func (sd *POCSensorData) checkData(s string) {
	// integrity check by checksum
	hs := sha256.New()
	io.WriteString(hs, s)
	calhs := fmt.Sprintf("%x", hs.Sum(nil))
	if calhs == sd.HashValue {
		fmt.Println("Input data checksum check OK")
	} else {
		fmt.Println("Input data checksum check error")
		os.Exit(11)
	}

	// delayed transmission check by timestamp
	ct := int(time.Now().Unix())
	itInt, _ := strconv.Atoi(sd.UnixTime)
	if ct-itInt/1000 < TimeThres {
		fmt.Println("Input data transmission on time OK")
	} else {
		fmt.Println("Input data delayed by ", ct-itInt/1000, "seconds")
		os.Exit(12)
	}
}
