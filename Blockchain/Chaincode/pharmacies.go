package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

func (c *Chaincode) GiveDrugs(ctx CustomTransactionContextInterface, drugID string, ignoredD, ignoredM []string) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Drugs with ID: %v doesn't exists", drugID)
	}
	var drugs Drugs
	json.Unmarshal(existing, &drugs)
	// check whether this pharmacies stores have roles as pharmacies
	if len(ignoredD) != len(ignoredM) {
		return Errorf("ignored drugs and ignored message missmacth")
	}
	if len(ignoredD) == 0 {
		drugs.Status = 2
	} else {
		for i, v := range ignoredD {
			if _, ok := drugs.Drug[v]; !ok {
				continue
			}
			drugs.Ignored[v] = ignoredM[i]
		}
		drugs.Status = 1
	}
	drugs.UpdateTime = time.Now().Unix()
	drugAsByte, _ := json.Marshal(drugs)

	return ctx.GetStub().PutState(drugs.ID, drugAsByte)
}
