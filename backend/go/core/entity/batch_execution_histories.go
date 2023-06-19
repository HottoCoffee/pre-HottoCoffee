package entity

import (
	"sort"
	"time"
)

type BatchExecutionHistories struct {
	Batch     Batch
	Histories []History
}

func NewBatchExecutionHistories(batch Batch, histories []History) BatchExecutionHistories {
	sort.Slice(histories, func(i, j int) bool {
		a := histories[i]
		b := histories[j]

		if !a.StartDatetime.Equal(b.StartDatetime) {
			return a.StartDatetime.Before(b.StartDatetime)
		} else {
			return a.FinishDatetime.Before(b.FinishDatetime)
		}
	})

	setTimeHistoriesMap := map[time.Time][]History{}
	var setDatetimes []time.Time
	for _, h := range histories {
		setDatetime := batch.CronSetting.Prev(h.StartDatetime.Add(time.Second))
		if setDatetime.Before(batch.StartDate) ||
			batch.EndDate != nil && setDatetime.After(*batch.EndDate) {
			continue
		}
		setTimeHistoriesMap[setDatetime] = append(setTimeHistoriesMap[setDatetime], h)
		setDatetimes = append(setDatetimes, setDatetime)
	}

	var distinctHistories []History
	for _, setDatetime := range setDatetimes {
		histories := setTimeHistoriesMap[setDatetime]
		distinctHistories = append(distinctHistories, histories[len(histories)-1])
	}

	return BatchExecutionHistories{Batch: batch, Histories: distinctHistories}
}
