package entity

import (
	"errors"
	"fmt"
	"time"
)

type History struct {
	Id               *int
	Batch            *Batch
	Status           Status
	ReportedDatetime *time.Time
}

func NewHistory(id *int, batch *Batch, statusValue string, reportedDatetime *time.Time) (*History, error) {
	status, err := mapToStatus(statusValue)
	if err != nil {
		return nil, err
	}

	if status == Success || status == Failed {
		if id == nil {
			return nil, errors.New("id should not be nil for success/failed history")
		}
		if batch == nil {
			return nil, errors.New("batch should not be nil for success/failed history")
		}
		if reportedDatetime != nil {
			return nil, errors.New("reportedDatetime should be nil for success/failed history")
		}
		return &History{
			Id:               id,
			Batch:            batch,
			Status:           status,
			ReportedDatetime: reportedDatetime,
		}, nil
	} else {
		if id != nil {
			return nil, errors.New("id should be nil for before started/in progress history")
		}
		if batch != nil {
			return nil, errors.New("batch should be nil for before started/in progress history")
		}
		if reportedDatetime != nil {
			return nil, errors.New("reportedDatetime should be nil for before started/in progress history")
		}
		return &History{Status: status}, nil
	}
}

type Status string

const (
	BeforeStarted Status = "before_started"
	InProgress    Status = "in_progress"
	Success       Status = "success"
	Failed        Status = "failed"
)

func mapToStatus(v string) (Status, error) {
	switch v {
	case "before_started":
		return BeforeStarted, nil
	case "in_progress":
		return InProgress, nil
	case "success":
		return Success, nil
	case "failed":
		return Failed, nil
	default:
		return "", errors.New(fmt.Sprintf("wrong history status value: %v", v))
	}
}

var BeforeStartedHistory History = History{
	Id:               nil,
	Status:           BeforeStarted,
	ReportedDatetime: nil,
}

var InProgressHistory History = History{
	Id:               nil,
	Status:           InProgress,
	ReportedDatetime: nil,
}
