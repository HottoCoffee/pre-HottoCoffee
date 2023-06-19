package core

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"time"
)

func GenerateCalendarItems(historiesArray []entity.BatchExecutionHistories, now time.Time, startDatetime time.Time, endDatetime time.Time) []entity.CalendarItem {
	var calendarItems []entity.CalendarItem
	for _, batchExecutionHistories := range historiesArray {
		logicalTimeHistoryMap := map[time.Time]entity.History{}
		batch := batchExecutionHistories.Batch
		for _, history := range batchExecutionHistories.Histories {
			logicalStartTime := batch.CronSetting.Prev(history.StartDatetime)
			logicalTimeHistoryMap[logicalStartTime] = history
		}

		begin := max(startDatetime, batch.StartDate)
		var end time.Time
		if batch.EndDate == nil {
			end = endDatetime
		} else {
			end = min(*batch.EndDate, endDatetime)
		}
		for _, logicalStartTime := range batch.CronSetting.ListSchedules(begin, end) {
			if history, ok := logicalTimeHistoryMap[logicalStartTime]; ok {
				calendarItems = append(calendarItems, entity.NewCalendarItemFromHistory(batch, history))
				continue
			}

			if logicalStartTime.After(now) {
				calendarItems = append(calendarItems, entity.NewNotExecutedCalendarItem(batch, entity.BeforeStated, logicalStartTime))
				continue
			}

			if (now.Equal(logicalStartTime) || now.After(logicalStartTime)) && now.Before(logicalStartTime.Add(time.Minute*time.Duration(batch.TimeLimit+1))) {
				calendarItems = append(calendarItems, entity.NewNotExecutedCalendarItem(batch, entity.InProgress, logicalStartTime))
				continue
			}

			calendarItems = append(calendarItems, entity.NewNotExecutedCalendarItem(batch, entity.FailedStatus, logicalStartTime))
		}
	}

	return calendarItems
}

func min(a time.Time, b time.Time) time.Time {
	if a.Before(b) {
		return a
	} else {
		return b
	}
}

func max(a time.Time, b time.Time) time.Time {
	if a.After(b) {
		return a
	} else {
		return b
	}
}
