package main

import (
	"encoding/json"
	. "fmt"
	"strconv"
	"time"
)

const (
	CREATED = "CREATED"
	UPDATED = "UPDATED"
)

type ReportOutput struct {
	Result []Report `json:"reports"`
}
type TreatmentOutput struct {
	Result []Treatment `json:"result"`
}
type TestOutput struct {
	Result []Test `json:"result"`
}

func (c *Chaincode) RefTest(ctx CustomTransactionContextInterface, reportID, name, refDoctor string, typeoftest int) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
	}

	var report Report
	json.Unmarshal(ctx.GetData(), &report)
	// if report.Doctor != doc from certs {
	// 	return Errorf("Cannot prescribe drug from the patient")
	// }
	if ok := c.checkConsent(ctx, report.PatientID, refDoctor); !ok {
		return "", Errorf("No consent from the patient")
	}

	id := TESTS + getSafeRandomString(ctx.GetStub())
	test := Test{
		DocTyp:            TESTS,
		ReportID:          reportID,
		ID:                id,
		Name:              name,
		RefDoctor:         refDoctor,
		Status:            0,
		MediaFileLocation: []string{},
		CreateTime:        time.Now().Unix(),
		UpdateTime:        time.Now().Unix(),
		PatientID:         report.PatientID,
	}
	if typeoftest == 1 {
		test.TypeOfT = 1
	}
	testAsByte, _ := json.Marshal((test))
	return test.ID, ctx.GetStub().PutState(id, testAsByte)
}
func (c *Chaincode) RefTreatment(ctx CustomTransactionContextInterface, reportID, refDoctor, name string) (string, error) {
	if ctx.GetData() == nil {
		return "", Errorf("Report with ID %v doesn't exists", reportID)
	}

	var report Report
	json.Unmarshal(ctx.GetData(), &report)
	// if report.Doctor != doc from certs {
	// 	return Errorf("Cannot prescribe drug from the patient")
	// }
	if ok := c.checkConsent(ctx, report.PatientID, refDoctor); !ok {
		return "", Errorf("No consent from the patient")
	}
	id := TREATMENT + getSafeRandomString(ctx.GetStub())
	treatment := Treatment{
		DocTyp:            TREATMENT,
		ReportID:          reportID,
		ID:                id,
		RefDoctor:         refDoctor,
		Name:              name,
		Comments:          make(map[string]string),
		Status:            0,
		MediaFileLocation: []string{},
		CreateTime:        time.Now().Unix(),
		UpdateTime:        time.Now().Unix(),
		PatientID:         report.PatientID,
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
	id := DRUGS + getSafeRandomString(ctx.GetStub())
	drugs := Drugs{
		DocTyp:     DRUGS,
		ReportID:   reportID,
		ID:         id,
		RefDoctor:  refDoctor,
		Drug:       make(map[string]string),
		Status:     0,
		Pending:    make(map[string]string),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	if len(drug) != len(doses) {
		return "", Errorf("Error: Missmatch with length of drug and doses")
	}
	for i, d := range drug {
		drugs.Drug[d] = doses[i]
		drugs.Pending[d] = "init"
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
	timeNow := time.Now().Unix()
	stringTime := strconv.FormatInt(timeNow, 10)
	report.UpdateTime = timeNow
	report.Comments[stringTime] = comment
	reportAsByte, _ := json.Marshal(report)

	return ctx.GetStub().PutState(report.ID, reportAsByte)
}

func (c *Chaincode) AddCommentsToTreatment(ctx CustomTransactionContextInterface, treatmentID, superviosr, comment string) error {
	if ctx.GetData() == nil {
		return Errorf("Treatment with ID %v doesn't exists", treatmentID)
	}
	var treatment Treatment
	json.Unmarshal(ctx.GetData(), &treatment)
	if ok := c.checkConsent(ctx, treatment.PatientID, superviosr); !ok {
		return Errorf("no consent from the patient")
	}
	if treatment.Status == 2 {
		return Errorf("Treatment is already completed")
	}
	timeNow := time.Now().Unix()
	stringTime := strconv.FormatInt(timeNow, 10)
	treatment.Comments[stringTime] = comment

	treatment.UpdateTime = timeNow
	treatmentAsByte, _ := json.Marshal(treatment)
	return ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}

func (c *Chaincode) AddMediaToTreatment(ctx CustomTransactionContextInterface, treatmentID, superviosr string, numberOfMfile int) ([]string, error) {
	if ctx.GetData() == nil {
		return []string{}, Errorf("Treatment with ID %v doesn't exists", treatmentID)
	}
	var treatment Treatment
	json.Unmarshal(ctx.GetData(), &treatment)
	if ok := c.checkConsent(ctx, treatment.PatientID, superviosr); !ok {
		return []string{}, Errorf("No consent from the patient")
	}
	for i := 0; i < numberOfMfile; i++ {
		id := TREATMENT + "media" + getSafeRandomString(ctx.GetStub()) + strconv.Itoa(i)
		treatment.MediaFileLocation = append(treatment.MediaFileLocation, id)
	}
	treatment.UpdateTime = time.Now().Unix()
	treatmentAsByte, _ := json.Marshal(treatment)
	return treatment.MediaFileLocation, ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}
