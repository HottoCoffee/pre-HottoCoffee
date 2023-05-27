package entity

import "fmt"

type BatchExecutionHistory struct {
	Batch   Batch
	History History
}

func NewBatchExecutionHistory(batch Batch, history History) (*BatchExecutionHistory, error) {
	if history.StartDatetime.Before(batch.StartDate) {
		return nil, NewDomainRuleViolationError(fmt.Sprintf("history should be started after batch's start date, batch id: %v, history id: %v, batch start datetime: %v, history start datetime: %v", batch.Id, history.Id, batch.StartDate, history.StartDatetime))
	}

	if batch.EndDate != nil && !history.StartDatetime.Before(batch.EndDate.AddDate(0, 0, 1)) {
		return nil, NewDomainRuleViolationError(fmt.Sprintf("history should not be started after batch's end date, batch id: %v, history id: %v, batch start datetime: %v, history start datetime: %v", batch.Id, history.Id, batch.StartDate, history.StartDatetime))
	}

	return &BatchExecutionHistory{Batch: batch, History: history}, nil
}
