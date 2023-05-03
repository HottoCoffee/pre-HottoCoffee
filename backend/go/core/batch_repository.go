package core

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"time"
)

type BatchRepository interface {
	FindById(id int) (*entity.Batch, error)
	FindAll() ([]entity.Batch, error)
	FindFilteredBy(query string) ([]entity.Batch, error)
	Create(batchName string, serverName string, cronSetting string, timeLimit int, startDate time.Time) (*entity.Batch, error)
}
