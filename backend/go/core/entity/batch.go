package entity

import (
	"errors"
	"strconv"
	"time"
)

type Batch struct {
	Id                int
	BatchName         string
	ServerName        string
	CronSetting       CronSetting
	TimeLimit         int
	EstimatedDuration int
	StartDate         time.Time
	EndDate           *time.Time
}

func NewBatch(
	id int,
	batchName string,
	serverName string,
	cronSetting string,
	timeLimit int,
	estimationDuration int,
	startDate time.Time,
	endDate *time.Time,
) (*Batch, error) {
	if id <= 0 {
		return nil, errors.New("ID should be more than 0. Given: " + strconv.Itoa(id))
	}
	if len(batchName) == 0 {
		return nil, errors.New("batch name should not be empty")
	}
	if len(serverName) == 0 {
		return nil, errors.New("server name should not be empty")
	}
	cs, err := NewCronSetting(cronSetting)
	if err != nil {
		return nil, err
	}
	if timeLimit < 1 {
		return nil, errors.New("time limit should be equal or more than 1. Given: " + strconv.Itoa(timeLimit))
	}
	if !(estimationDuration >= 0 && estimationDuration < timeLimit) {
		return nil, errors.New("estimation duration should be equal or more than 0 and less than time limit")
	}
	if endDate != nil && (endDate.Equal(startDate) || endDate.Before(startDate)) {
		return nil, errors.New("end date should be after start date if exists")
	}

	return &Batch{
		Id:                id,
		BatchName:         batchName,
		ServerName:        serverName,
		CronSetting:       *cs,
		TimeLimit:         timeLimit,
		EstimatedDuration: estimationDuration,
		StartDate:         startDate,
		EndDate:           endDate,
	}, nil
}