package entity

import (
	"time"
)

type CalendarItem struct {
	Batch          Batch
	Id             *int
	Status         CalendarStatus
	StartDatetime  time.Time
	FinishDatetime *time.Time
}

func NewCalendarItemFromHistory(batch Batch, history History) CalendarItem {
	return CalendarItem{
		Batch:          batch,
		Id:             &history.Id,
		Status:         mapExecutionResultToCalendarStatus(history.ExecutionResult),
		StartDatetime:  history.StartDatetime,
		FinishDatetime: &history.FinishDatetime,
	}
}

func NewNotExecutedCalendarItem(batch Batch, status CalendarStatus, startDatetime time.Time) CalendarItem {
	return CalendarItem{
		Batch:          batch,
		Id:             nil,
		Status:         status,
		StartDatetime:  startDatetime,
		FinishDatetime: nil,
	}
}

type CalendarStatus string

const (
	SuccessStatus CalendarStatus = "success"
	FailedStatus  CalendarStatus = "failed"
	InProgress    CalendarStatus = "in_progress"
	BeforeStated  CalendarStatus = "before_started"
)

func mapExecutionResultToCalendarStatus(result ExecutionResult) CalendarStatus {
	if result == Success {
		return SuccessStatus
	} else {
		return FailedStatus
	}
}
