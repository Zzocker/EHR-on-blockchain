package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func getSafeRandomString(stub shim.ChaincodeStubInterface) string {
	// TODO : use timestamp and txID as seed for pseudo number generator to use uuid package
	return stub.GetTxID()
}
