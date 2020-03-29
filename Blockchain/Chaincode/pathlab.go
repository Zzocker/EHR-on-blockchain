package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
)

type OutputResult struct {
	MediaFile []string `josn:"media_file"`
	Type      int      `json:"type_of_test"`
}

func (c *Chaincode) DoTest(ctx CustomTransactionContextInterface, testID, result, supervisor string, numberOfMfile int) (OutputResult, error) {
	existing := ctx.GetData()
	if existing == nil {
		return OutputResult{}, Errorf("test with ID: %v doesn't exists", testID)
	}
	var test Test
	json.Unmarshal(existing, &test)
	if test.Status == 1 {
		return OutputResult{}, Errorf("test is already done")
	}
	test.UpdateTime = time.Now().Unix()
	for i := 0; i < numberOfMfile; i++ {
		test.MediaFileLocation = append(test.MediaFileLocation, uuid.New().String())
	}
	test.Supervisor = supervisor
	test.Status = 1
	test.Result = result
	testAsByte, _ := json.Marshal(test)

	return OutputResult{MediaFile: test.MediaFileLocation, Type: test.TypeOfT}, ctx.GetStub().PutState(test.ID, testAsByte)
}