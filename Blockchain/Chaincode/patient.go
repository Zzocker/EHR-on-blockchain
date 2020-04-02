package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

func (c *Chaincode) RegisterPatient(ctx CustomTransactionContextInterface, aadhaar, permCon string) error {
	existing := ctx.GetData()
	if existing != nil {
		return Errorf("Aadhaar ID allready exists")
	}
	consent := Consent{
		DocTyp:              CONSENT,
		ID:                  aadhaar,
		PermanentConsenters: make(map[string]bool),
		TemporaryConsenters: make(map[string]int64),
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

func (c *Chaincode) GetTest(ctx CustomTransactionContextInterface, key, requester string) (Test, error) {
	existing := ctx.GetData()
	if existing == nil {
		return Test{}, Errorf("Test with ID %v does'nt exists", key)
	}

	var test Test
	json.Unmarshal(existing, &test)
	if ok := c.checkConsent(ctx, test.PatientID, requester); !ok && test.TypeOfT == 0 {
		return Test{}, Errorf("Please get consent form the Patient")
	}
	return test, nil
}
func (c *Chaincode) GetReport(ctx CustomTransactionContextInterface, key, requester string) (Report, error) {
	existing := ctx.GetData()
	if existing == nil {
		return Report{}, Errorf("Report with ID %v does'nt exists", key)
	}
	var report Report
	json.Unmarshal(existing, &report)
	if ok := c.checkConsent(ctx, report.PatientID, requester); !ok {
		return Report{}, Errorf("Please get consent form the Patient")
	}
	return report, nil
}
func (c *Chaincode) GetTreatment(ctx CustomTransactionContextInterface, key, requester string) (Treatment, error) {
	existing := ctx.GetData()
	if existing == nil {
		return Treatment{}, Errorf("Treatment with ID %v does'nt exists", key)
	}
	var treatment Treatment
	json.Unmarshal(existing, &treatment)
	if ok := c.checkConsent(ctx, treatment.PatientID, requester); !ok {
		return Treatment{}, Errorf("Please get consent form the Patient")
	}
	return treatment, nil
}

func (c *Chaincode) GetDrugs(ctx CustomTransactionContextInterface, key, requester string) (Drugs, error) {
	existing := ctx.GetData()
	if existing == nil {
		return Drugs{}, Errorf("Treatment with ID %v does'nt exists", key)
	}
	var drugs Drugs
	json.Unmarshal(existing, &drugs)
	if ok := c.checkConsent(ctx, drugs.For, requester); !ok {
		return Drugs{}, Errorf("Please get consent form the Patient")
	}
	return drugs, nil
}