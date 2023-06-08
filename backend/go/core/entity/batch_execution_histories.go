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
	for _, h := range histories {
		setDatetime := batch.CronSetting.Prev(h.StartDatetime)
		if setDatetime.Before(batch.StartDate) ||
			batch.EndDate != nil && setDatetime.After(*batch.EndDate) {
			continue
		}
		setTimeHistoriesMap[setDatetime] = append(setTimeHistoriesMap[setDatetime], h)
	}

	var distinctHistories []History
	for _, v := range setTimeHistoriesMap {
		distinctHistories = append(distinctHistories, v[len(v)-1])
	}

	return BatchExecutionHistories{Batch: batch, Histories: distinctHistories}
}
