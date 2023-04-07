package models

type Outcome struct {
	Problem string //may include OperationSetId in the future
	Result string
	Reason string
}

type PreviousOutcomes struct {
	Outcomes map[string]Outcome
}