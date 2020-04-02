package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

func (c *Chaincode) GiveDrugs(ctx CustomTransactionContextInterface, drugID string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Drugs with ID: %v doesn't exists", drugID)
	}
	var drugs Drugs
	json.Unmarshal(existing, &drugs)
	if drugs.Status == 1{
		return Errorf("Drugs allready given")
	}
	// check whether this pharmacies stores have roles as pharmacies
	drugs.Status = 1
	drugs.UpdateTime = time.Now().Unix()
	drugAsByte, _ := json.Marshal(drugs)

	return ctx.GetStub().PutState(drugs.ID, drugAsByte)
}
