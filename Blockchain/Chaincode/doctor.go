package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

const (
	CREATED = "CREATED"
	UPDATED = "UPDATED"
)

func (c *Chaincode) RefTest(ctx CustomTransactionContextInterface, reportID, patientID, name, refDoctor string, typeoftest int) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
	}
	if ok := c.checkConsent(ctx, patientID, refDoctor); !ok {
		return "", Errorf("No consent from the patient")
	}
	id := uuid.New().String()
	test := Test{
		DocTyp:     TESTS,
		ReportID:   reportID,
		ID:         id,
		Name:       name,
		RefDoctor:  refDoctor,
		Status:     0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		PatientID:  patientID,
	}
	if typeoftest == 1 {
		test.TypeOfT = 1
	}
	testAsByte, _ := json.Marshal((test))
	return test.ID, ctx.GetStub().PutState(id, testAsByte)
}
func (c *Chaincode) RefTreatment(ctx CustomTransactionContextInterface, reportID, patientID, refDoctor, name string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
	}
	if ok := c.checkConsent(ctx, patientID, refDoctor); !ok {
		return "", Errorf("No consent from the patient")
	}
	id := uuid.New().String()
	treatment := Treatment{
		DocTyp:     TREATMENT,
		ReportID:   reportID,
		ID:         id,
		RefDoctor:  refDoctor,
		Name:       name,
		Comments:   make(map[int64]string),
		Status:     0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		PatientID:  patientID,
	}
	treatmentAsByte, _ := json.Marshal(treatment)
	return treatment.ID, ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}

func (c *Chaincode) PrescribeDrugs(ctx CustomTransactionContextInterface, reportID, refDoctor string, drug, doses []string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
	}

	var report Report
	json.Unmarshal(ctx.GetData(), &report)
	if ok := c.checkConsent(ctx, report.PatientID, refDoctor); !ok {
		return "", Errorf("No consent from the patient")
	}
	id := uuid.New().String()
	drugs := Drugs{
		DocTyp:     DRUGS,
		ReportID:   reportID,
		ID:         id,
		RefDoctor:  refDoctor,
		Drug:       make(map[string]string),
		Status:     0,
		Ignored:    make(map[string]string),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	if len(drug) != len(doses) {
		return "", Errorf("Error: Missmatch with length of drug and doses")
	}
	for i, d := range drug {
		drugs.Drug[d] = doses[i]
	}
	drugsAsByte, _ := json.Marshal(drugs)

	return drugs.ID, ctx.GetStub().PutState(id, drugsAsByte)
}

func (c *Chaincode) AddCommentsToReport(ctx CustomTransactionContextInterface, reportID, comment, refDoctor string) error {
	if ctx.GetData() == nil {
		return Errorf("Report with ID %v doesn't exists", reportID)
	}

	var report Report
	json.Unmarshal(ctx.GetData(), &report)
	// if report.Doctor != doc from certs {
	// 	return Errorf("Cannot prescribe drug from the patient")
	// }
	if ok := c.checkConsent(ctx, report.PatientID, refDoctor); !ok {
		return Errorf("No consent from the patient")
	}
	report.UpdateTime = time.Now().Unix()
	report.Comments[time.Now().Unix()] = comment
	reportAsByte, _ := json.Marshal(report)

	return ctx.GetStub().PutState(report.ID, reportAsByte)
}

func (c *Chaincode) AddCommentsToTreatment(ctx CustomTransactionContextInterface, treatmentID, superviosr, comment string) error {
	if ctx.GetData() == nil {
		return Errorf("Report with ID %v doesn't exists", treatmentID)
	}
	var treatment Treatment
	json.Unmarshal(ctx.GetData(), &treatment)
	if ok := c.checkConsent(ctx, treatment.PatientID, superviosr); !ok {
		return Errorf("No consent from the patient")
	}
	if treatment.Status == 2 {
		Errorf("Treatment is already completed")
	}
	treatment.Comments[time.Now().Unix()] = comment

	treatment.UpdateTime = time.Now().Unix()
	treatmentAsByte, _ := json.Marshal(treatment)
	return ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}

func (c *Chaincode) GetReports(ctx CustomTransactionContextInterface, requester, queryS, qtype string) ([]Report, error) {
	var output []Report
	var query string
	if qtype == CREATED {
		query = `{"use_index": "OnCreatedTime",`
	} else if qtype == UPDATED {
		query = `{"use_index": "OnUpdatedTime",`
	} else {
		return []Report{}, Errorf("Error : No such query type %v for lead", qtype)
	}
	query += `
			"selector": {
				"docTyp": "REPORT"`
	query += queryS
	// worldState, _ := ctx.GetStub().GetState(WORLDSTATE)
	// to add into selector , ---new selector----}}
	// not to selector , --},new :----}
	result, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return []Report{}, err
	}
	for result.HasNext() {
		var resultKV *queryresult.KV

		resultKV, _ = result.Next()
		var report Report
		json.Unmarshal(resultKV.GetValue(), &report)
		if ok := c.checkConsent(ctx, report.PatientID, requester); !ok {
			continue
		}
		output = append(output, report)
	}
	return output, result.Close()
}

func (c *Chaincode) GetTreatment(ctx CustomTransactionContextInterface, requester, queryS, qtype string) ([]Treatment, error) {
	var output []Treatment
	var query string
	if qtype == CREATED {
		query = `{"use_index": "OnCreatedTime",`
	} else if qtype == UPDATED {
		query = `{"use_index": "OnUpdatedTime",`
	} else {
		return []Treatment{}, Errorf("Error : No such query type %v for lead", qtype)
	}
	query += `
			"selector": {
				"docTyp": "TREATMENT"`
	query += queryS
	// worldState, _ := ctx.GetStub().GetState(WORLDSTATE)
	// to add into selector , ---new selector----}}
	// not to selector , --},new :----}
	result, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return []Treatment{}, err
	}
	for result.HasNext() {
		var resultKV *queryresult.KV

		resultKV, _ = result.Next()
		var treatment Treatment
		json.Unmarshal(resultKV.GetValue(), &treatment)
		if ok := c.checkConsent(ctx, treatment.PatientID, requester); !ok {
			continue
		}
		output = append(output, treatment)
	}
	return output, result.Close()
}
