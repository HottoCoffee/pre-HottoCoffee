package core

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type HistoryRepository interface {
	FindByHistoryIdAndBatchId(historyId int, batchId int) (*entity.BatchExecutionHistory, error)
}
