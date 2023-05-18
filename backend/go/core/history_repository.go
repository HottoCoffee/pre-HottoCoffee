package core

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"time"
)

type HistoryRepository interface {
	FindByHistoryIdAndBatchId(historyId int, batchId int) (*entity.History, error)
	FindAllByBatchId(batchId int) (*entity.Histories, error)
	FindAllDuring(start time.Time, endInclusive time.Time) ([]entity.Histories, error)
}
