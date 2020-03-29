package main

const WORLDSTATE = "WORLDSTATE"

//ChangeState 0 - normal and 1 - pandemic
func ChangeState(ctx CustomTransactionContextInterface, newState int) error {
	return ctx.GetStub().PutState(WORLDSTATE, newState)
}
