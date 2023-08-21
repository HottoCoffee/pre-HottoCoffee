package core

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"time"
)

type HistoryRepository interface {
	FindByHistoryIdAndBatchId(historyId int, batchId int) (*entity.BatchExecutionHistory, error)
	FindByBatchId(batchId int) (*entity.BatchExecutionHistories, error)
	FindAllDuring(startDate time.Time, endDate time.Time) ([]entity.BatchExecutionHistories, error)
}
