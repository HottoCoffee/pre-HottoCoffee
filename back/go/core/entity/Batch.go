package entity

import (
	"errors"
	"time"
)

type Batch struct {
	Id                 int
	BatchName          string
	ServerName         string
	CronSetting        string // TODO: use CronSetting-dedicated class
	TimeLimit          int
	EsitimatedDuration int
	StartDate          time.Time
	EndDate            *time.Time
}

func NewBatch(id int) (*Batch, error) {
	if id <= 0 {
		return nil, errors.New("dddd")
	}
	return &Batch{}, nil
}
