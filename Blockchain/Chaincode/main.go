package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


// Chaincode : strcut for implementing chaincode
type Chaincode struct {
	contractapi.Contract
}

func main() {
	crmContract := new(Chaincode)
	crmContract.TransactionContextHandler = new(CustomTransactionContext)
	crmContract.BeforeTransaction = GetWorldState

	cc, err := contractapi.NewChaincode(crmContract)

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}
