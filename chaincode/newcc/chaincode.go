/*
 * SPDX-License-Identifier: Apache-2.0
 * UNDER THE AUTHORITY OF CHUNG YUP KIM & ECS VUW
 */

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

//	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric-chaincode-go/shim"
//	sc "github.com/hyperledger/fabric/protos/peer"
	sc "github.com/hyperledger/fabric-protos-go/peer"
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

// POCActData is the definition of the state database structure.
type POCActData struct {
	SensorID   string `json:"sid"`
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
	case "setData":
		return cc.setData(stub, params)
	case "queryData":
		return cc.queryData(stub, params)
	case "updateData":
		return cc.updateData(stub, params)
	default:
		return shim.Error("Input function name error")
	}
}

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

// QUERY : Query actuator status //
// Params : SensorID //
// Return : Actuator status //
func (cc *POCChaincode) queryData(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	// number of params check
	if len(params) != 1 {
		return shim.Error("Incorrect no. of parameters (queryData)")
	}
	sensorID := params[0]

	// query status
	actuatorStatus, err := stub.GetState(sensorID)
	if err != nil {
		return shim.Error("failed to query status, error: " + err.Error())
	}
	if actuatorStatus == nil {
		return shim.Error("failed to query status which is null")
	}

	return shim.Success(actuatorStatus)
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
	actuatorBytes, err := stub.GetState(sensorID)
	if err != nil {
		return shim.Error("failed to get queryData for updateData, error: " + err.Error())
	}
	actuator := POCActData{}
	err = json.Unmarshal(actuatorBytes, &actuator)
	if err != nil {
		return shim.Error("failed to unmarshal for updateData, error: " + err.Error())
	}

	// check threshold & status
	switch {
	case tempInt > MaxThres && actuator.Status == ActOn:
		return shim.Success([]byte("Actuator already ON"))
	case tempInt > MaxThres && actuator.Status == ActOff:
		actuator.Status = ActOn
		actuatorBytes, _ := json.Marshal(actuator)
		err = stub.PutState(sensorID, actuatorBytes)
		if err != nil {
			return shim.Error("failed to update status to ON : " + err.Error())
		}
		err = stub.SetEvent("updateEvent", actuatorBytes)
		if err != nil {
			return shim.Error("failed to emit event: " + err.Error())
		}
	case tempInt < MinThres && actuator.Status == ActOn:
		actuator.Status = ActOff
		actuatorBytes, _ := json.Marshal(actuator)
		err = stub.PutState(sensorID, actuatorBytes)
		if err != nil {
			return shim.Error("failed to update status to OFF : " + err.Error())
		}
		err = stub.SetEvent("updateEvent", actuatorBytes)
		if err != nil {
			return shim.Error("failed to emit event: " + err.Error())
		}
	case tempInt < MinThres && actuator.Status == ActOff:
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

func main() {
	if err := shim.Start(new(POCChaincode)); err != nil {
		fmt.Printf("Error starting POCChaincode chaincode: %s", err)
	}
}
