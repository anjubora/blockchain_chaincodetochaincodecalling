/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("example_cc2")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	logger.Info("########### example_cc2 Init ###########")

	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	id, err := cid.GetID(stub)
	fmt.Println("id is", id)

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	logger.Info("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "hello" {
		// Deletes an entity from its state
		return t.hello(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) hello(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var C string = args[0]
	var a = args[1]
	var X int
	var err error

	X, err = strconv.Atoi(a)
	// must be an invoke
	// var A, B string    // Entities
	// var Aval, Bval int // Asset holdings
	//         // Transaction value
	//
	// if len(args) != 3 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
	// }

	// A = args[0]
	// B = args[1]

	// // Get the state from the ledger
	// // TODO: will be nice to have a GetAllState call to ledger
	// Avalbytes, err := stub.GetState(A)
	// if err != nil {
	// 	return shim.Error("Failed to get state")
	// }
	// if Avalbytes == nil {
	// 	return shim.Error("Entity not found")
	// }
	// Aval, _ = strconv.Atoi(string(Avalbytes))

	// Bvalbytes, err := stub.GetState(B)
	// if err != nil {
	// 	return shim.Error("Failed to get state")
	// }
	// if Bvalbytes == nil {
	// 	return shim.Error("Entity not found")
	// }
	// Bval, _ = strconv.Atoi(string(Bvalbytes))

	// // Perform the execution
	// X, err = strconv.Atoi(args[2])
	// if err != nil {
	// 	return shim.Error("Invalid transaction amount, expecting a integer value")
	// }
	// Aval = Aval - X
	// Bval = Bval + X
	logger.Infof("hello %s", a)

	// // Write the state back to the ledger
	err = stub.PutState(C, []byte(strconv.Itoa(X)))
	if err != nil {
		return shim.Error(err.Error())
	}

	// err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
