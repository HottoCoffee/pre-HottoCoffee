package core

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type HistoryRepository interface {
	FindByIdAndBatchId(historyId int, batchId int) (*entity.History, error)
}
