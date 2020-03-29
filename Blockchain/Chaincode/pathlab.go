package main
import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
)

func (c *Chaincode) DoTest(ctx CustomTransactionContextInterface,testID,supervisor string,numberOfMfile int)([]string,error) {
	existing:= ctx.GetData()
	if existing==nil{
		return []string{},Errorf("test with ID: %v doesn't exists",testID)
	}
	var test Test
	json.Unmarshal(existing,&test)
	if test.Status == 1 {
		return []string{},Errorf("test is already done")
	}
	mfileID = uuid.New().String()
	test.UpdateTime=time.Now().Unix()
	for i:=0;i<numberOfMfile;i++{
		test.MediaFileLocation = append(test.MediaFileLocation,uuid.New().String())
	}
	test.Supervisor = supervisor
	test.Status=1
	testAsByte,_:= json.Marshal(test)

	return test.MediaFileLocation,ctx.GetStub().PutState(test.ID,testAsByte)
}
