package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
)



func (c *Chaincode) RefTest(ctx CustomTransactionContextInterface, reportID, name, refDoctor string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
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
	}
	testAsByte, _ := json.Marshal((test))
	return test.ID, ctx.GetStub().PutState(id, testAsByte)
}
func (c *Chaincode) RefTreatment(ctx CustomTransactionContextInterface, reportID, refDoctor, name string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
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
	}
	treatmentAsByte, _ := json.Marshal(treatment)
	return treatment.ID, ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}

func (c *Chaincode) PrescribeDrugs(ctx CustomTransactionContextInterface, reportID, patientID, refDoctor string, drug, doses []string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
	}
	var report Report
	json.Unmarshal(ctx.GetData(), &report)
	if report.PatientID != patientID {
		return "", Errorf("Cannot prescribe drug for this patient")
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
	if len(drug)!=len(doses){
		return "",Errorf("Error: Missmatch with length of drug and doses")
	}
	for i, d := range drug {
		drugs.Drug[d] = doses[i]
	}
	drugsAsByte, _ := json.Marshal(drugs)

	return drugs.ID, ctx.GetStub().PutState(id, drugsAsByte)
}

func (c *Chaincode) AddComments(ctx CustomTransactionContextInterface, reportID, comment string) error {
	if ctx.GetData() == nil {
		return Errorf("Report with ID %v doesn't exists", reportID)
	}
	var report Report
	json.Unmarshal(ctx.GetData(), &report)
	// if report.Doctor != doc from certs {
	// 	return Errorf("Cannot prescribe drug for this patient")
	// }
	report.UpdateTime = time.Now().Unix()
	report.Comments[time.Now().Unix()] = comment
	reportAsByte, _ := json.Marshal(report)

	return ctx.GetStub().PutState(report.ID, reportAsByte)
}

func (c *Chaincode) DoTreatment(ctx CustomTransactionContextInterface,treatmentID,supervisor string)  {
	
}