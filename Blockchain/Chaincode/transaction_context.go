package main

import (
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//CustomTransactionContext : transaction
type CustomTransactionContext struct {
	contractapi.TransactionContext
	data []byte
}

// GetData : return the date stored in ctc
func (ctc *CustomTransactionContext) GetData() []byte {
	return ctc.data
}

// SetData : sets the data to the ctc
func (ctc *CustomTransactionContext) SetData(data []byte) {
	ctc.data = data
}

// CustomTransactionContextInterface : interface implementing get ctc
type CustomTransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetData() []byte
	SetData([]byte)
}

// GetWorldState : gets the world state and set it to ctc.data
func GetWorldState(ctx CustomTransactionContextInterface) error {
	_, params := ctx.GetStub().GetFunctionAndParameters()

	if len(params) < 1 {
		return errors.New("Missing key for world state")
	}

	existing, err := ctx.GetStub().GetState(params[0])

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	ctx.SetData(existing)

	return nil
}