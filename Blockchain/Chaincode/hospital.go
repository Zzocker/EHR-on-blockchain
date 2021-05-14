package main

import (
	"encoding/json"
	. "fmt"
	"time"
)

func (c *Chaincode) CreateNewReport(ctx CustomTransactionContextInterface, patientID, refDoctor string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Patient of ID %v doesn't exists", patientID)
	}
	id := REPORT + getSafeRandomString(ctx.GetStub())
	report := Report{
		DocTyp:      REPORT,
		ID:          id,
		PatientID:   patientID,
		Status:      "0",
		RefDoctorID: refDoctor,
		Comments:    make(map[string]string),
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}
	reportAsByte, _ := json.Marshal(report)
	return report.ID, ctx.GetStub().PutState(id, reportAsByte)
}

func (c *Chaincode) StartTreatment(ctx CustomTransactionContextInterface, treatmentID, supervisor string) error {
	if ctx.GetData() == nil {
		return Errorf("Treatment with ID %v doesn't exists", treatmentID)
	}
	var treatment Treatment
	json.Unmarshal(ctx.GetData(), &treatment)
	if treatment.Status != 0 {
		return Errorf("Cannot start allready started treatment")
	}
	treatment.Supervisor = supervisor
	treatment.Status = 1
	treatment.UpdateTime = time.Now().Unix()

	treatmentAsByte, _ := json.Marshal(treatment)

	return ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}
