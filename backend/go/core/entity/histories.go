package entity

import (
	"errors"
	"sort"
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
	sort.Slice(simpleHistories, func(i, j int) bool {
		return simpleHistories[i].ReportedDatetime.Before(simpleHistories[j].ReportedDatetime)
	})
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

func (hh Histories) GetHistoryExecutedOn(startDatetime time.Time) History {
	estimatedEndDatetime := startDatetime.Add(time.Duration(int(time.Minute) * hh.Batch.TimeLimit))
	var idx int
	if len(hh.SimpleHistories) == 0 {
		idx = -1
	} else {
		idx = hh.binarySearch(0, len(hh.SimpleHistories), startDatetime, estimatedEndDatetime)
	}
	if idx >= 0 {
		history, _ := NewHistory(
			&hh.SimpleHistories[idx].Id,
			&hh.Batch,
			string(hh.SimpleHistories[idx].Status),
			startDatetime,
			&hh.SimpleHistories[idx].ReportedDatetime,
		)
		return *history
	}

	now := time.Now()
	if isDuring(now, startDatetime, estimatedEndDatetime) {
		history, _ := NewHistory(nil, &hh.Batch, string(InProgress), startDatetime, nil)
		return *history
	} else if now.After(estimatedEndDatetime) {
		history, _ := NewHistory(nil, &hh.Batch, string(BeforeStarted), startDatetime, nil)
		return *history
	} else {
		history, _ := NewHistory(nil, &hh.Batch, string(Failed), startDatetime, nil)
		return *history
	}
}

func (hh Histories) binarySearch(beginIdx int, endIdx int, startDatetime time.Time, endDatetime time.Time) int {
	if endIdx-beginIdx <= 1 {
		history := hh.SimpleHistories[beginIdx]
		if isDuring(history.ReportedDatetime, startDatetime, endDatetime) {
			return beginIdx
		} else {
			return -1
		}
	}

	midIdx := (beginIdx + endIdx) / 2
	history := hh.SimpleHistories[midIdx]
	if isDuring(history.ReportedDatetime, startDatetime, endDatetime) {
		return midIdx
	} else if endDatetime.Before(history.ReportedDatetime) {
		return hh.binarySearch(beginIdx, midIdx, startDatetime, endDatetime)
	} else {
		return hh.binarySearch(midIdx, endIdx, startDatetime, endDatetime)
	}
}

func isDuring(target time.Time, begin time.Time, end time.Time) bool {
	isEqualOrAfterBegin := target.Equal(begin) || target.After(begin)
	isEqualOrBeforeEnd := target.Before(end) || target.Equal(end)
	return isEqualOrAfterBegin && isEqualOrBeforeEnd
}
