package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

func (c *Chaincode) RegisterPatient(ctx CustomTransactionContextInterface, aadhaar, permCon string) error {
	consent := Consent{
		DocTyp:              CONSENT,
		ID:                  aadhaar,
		PermanentConsenters: make(map[string]bool),
		TemporaryConsenters: make(map[string]int64),
		Status:              "Normal",
	}
	consent.PermanentConsenters[aadhaar] = true
	consent.PermanentConsenters[permCon] = true

	consentAsByte, _ := json.Marshal(consent)

	return ctx.GetStub().PutState(aadhaar, consentAsByte)
}

func (c *Chaincode) UpdateTempConsent(ctx CustomTransactionContextInterface, consentID, typeOfUpdate, to string, till int64) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Consent with key %v doesn't exists")
	}
	var consent Consent
	json.Unmarshal(existing, &consent)
	if typeOfUpdate == "ADD" {
		consent.TemporaryConsenters[to] = till
	} else if typeOfUpdate == "REMOVE" {
		if _, ok := consent.TemporaryConsenters[to]; ok {
			delete(consent.TemporaryConsenters, to)
		}
	}
	consentAsByte, _ := json.Marshal(consent)

	return ctx.GetStub().PutState(consent.ID, consentAsByte)
}

func (c *Chaincode) UpdatePermConsent(ctx CustomTransactionContextInterface, consentID, typeOfUpdate, to string, till int64) error {
	existing := ctx.GetData()
	if existing == nil {
		return Errorf("Consent with key %v doesn't exists")
	}
	var consent Consent
	json.Unmarshal(existing, &consent)
	if typeOfUpdate == "ADD" {
		consent.PermanentConsenters[to] = true
	} else if typeOfUpdate == "REMOVE" {
		if _, ok := consent.PermanentConsenters[to]; ok {
			delete(consent.PermanentConsenters, to)
		}
	}
	consentAsByte, _ := json.Marshal(consent)

	return ctx.GetStub().PutState(consent.ID, consentAsByte)
}

func (c *Chaincode) getByte(ctx CustomTransactionContextInterface, key string) ([]byte, error) {
	AsByte, _ := ctx.GetStub().GetState(key)
	if AsByte == nil {
		return []byte{}, Errorf("No state with key %v", key)
	}
	return AsByte, nil
}

func (c *Chaincode) checkConsent(ctx CustomTransactionContextInterface, consentID, checkFor string) bool {
	consentAsyte, err := c.getByte(ctx, consentID)
	if err != nil {
		return false
	}
	var consent Consent
	json.Unmarshal(consentAsyte, &consent)

	for k := range consent.PermanentConsenters {
		if checkFor == k {
			return true
		}
	}
	for k, v := range consent.TemporaryConsenters {
		if checkFor == k && v >= time.Now().Unix() {
			return true
		}
	}
	return false
}

func (c *Chaincode) GetTest(ctx CustomTransactionContextInterface, requester, queryS, qtype string) (TestOutput, error) {
	var output []Test
	var query string
	if qtype == CREATED {
		query = `{"use_index": "OnCreatedTime",`
	} else if qtype == UPDATED {
		query = `{"use_index": "OnUpdatedTime",`
	} else {
		return TestOutput{}, Errorf("Error : No such query type %v for lead", qtype)
	}
	query += `
			"selector": {
				"docTyp": "TESTS"`
	query += queryS
	// to add into selector , ---new selector----}}
	// not to selector , --},new :----}
	result, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return TestOutput{}, err
	}
	for result.HasNext() {
		var resultKV *queryresult.KV

		resultKV, _ = result.Next()
		var test Test
		json.Unmarshal(resultKV.GetValue(), &test)
		if ok := c.checkConsent(ctx, test.PatientID, requester); !ok {
			continue
		}
		if test.TypeOfT != 1 {
			continue
		}
		output = append(output, test)
	}
	return TestOutput{Result: output}, result.Close()
}

func (c *Chaincode) GetStateAsyte(ctx CustomTransactionContextInterface, key string) ([]byte, error) {
	return c.getByte(ctx, key)
}
