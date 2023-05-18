package entity

import (
	"errors"
	"fmt"
	"time"
)

type History struct {
	Id             *int
	Batch          *Batch
	Status         Status
	StartDatetime  *time.Time
	FinishDatetime *time.Time
}

func NewHistory(id *int, batch *Batch, statusValue string, startDatetime time.Time, reportedDatetime *time.Time) (*History, error) {
	status, err := mapToStatus(statusValue)
	if err != nil {
		return nil, err
	}

	if status == Success {
		if id == nil {
			return nil, errors.New("id should not be nil for success history")
		}
		if batch == nil {
			return nil, errors.New("batch should not be nil for success history")
		}
		if reportedDatetime == nil {
			return nil, errors.New("reportedDatetime should be nil for success history")
		}
		return &History{
			Id:             id,
			Batch:          batch,
			Status:         status,
			StartDatetime:  &startDatetime,
			FinishDatetime: reportedDatetime,
		}, nil
	} else if status == Failed {
		return &History{
			Id:             id,
			Batch:          batch,
			Status:         status,
			StartDatetime:  &startDatetime,
			FinishDatetime: reportedDatetime,
		}, err
	} else {
		if id != nil {
			return nil, errors.New("id should be nil for before started/in progress history")
		}
		if reportedDatetime != nil {
			return nil, errors.New("reportedDatetime should be nil for before started/in progress history")
		}
		return &History{
			Id:             id,
			Batch:          batch,
			Status:         status,
			StartDatetime:  &startDatetime,
			FinishDatetime: reportedDatetime,
		}, nil
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
