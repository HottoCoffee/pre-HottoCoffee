package entity

import (
	"errors"
	"time"
)

type Histories struct {
	Batch           Batch
	SimpleHistories []SimpleHistory
}

type SimpleHistory struct {
	Id               int
	Status           Status
	ReportedDatetime time.Time
}

func NewHistories(batch Batch, simpleHistories []SimpleHistory) Histories {
	return Histories{
		Batch:           batch,
		SimpleHistories: simpleHistories,
	}
}

func NewSimpleHistory(id int, statusValue string, reportedDatetime time.Time) (*SimpleHistory, error) {
	if id <= 0 {
		return nil, errors.New("history id should be equal or more than 1")
	}
	status, err := mapToStatus(statusValue)
	if err != nil {
		return nil, err
	}
	if !(status == Success || status == Failed) {
		return nil, errors.New("status of stored history should be success or failed")
	}
	return &SimpleHistory{
		Id:               id,
		Status:           status,
		ReportedDatetime: reportedDatetime,
	}, nil
}
