package entity

import (
	"fmt"
	"time"
)

type History struct {
	Id              int
	ExecutionResult ExecutionResult
	StartDatetime   time.Time
	FinishDatetime  time.Time
}

func NewHistory(id int, executionResultString string, startDatetime time.Time, finishDatetime time.Time) (*History, error) {
	if id <= 0 {
		return nil, NewDomainRuleViolationError(fmt.Sprintf("history id should be equal or greater than 0. given: %v", id))
	}

	executionResult, err := mapStringToExecutionResult(executionResultString)
	if err != nil {
		return nil, err
	}

	if finishDatetime.Before(startDatetime) {
		return nil, NewDomainRuleViolationError(fmt.Sprintf("history finish datetime should be after start datetime. start datetime: %v, finish datetime: %v", startDatetime, finishDatetime))
	}

	return &History{
		Id:              id,
		ExecutionResult: executionResult,
		StartDatetime:   startDatetime,
		FinishDatetime:  finishDatetime,
	}, nil
}

type ExecutionResult string

const (
	success ExecutionResult = "success"
	failed  ExecutionResult = "failed"
)

func mapStringToExecutionResult(value string) (ExecutionResult, error) {
	switch value {
	case "success":
		return success, nil
	case "failed":
		return failed, nil
	default:
		return "", NewDomainRuleViolationError(fmt.Sprintf("execution result should be success or failed. given: %v", value))
	}
}
