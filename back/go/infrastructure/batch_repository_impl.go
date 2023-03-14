package infrastructure

import (
	"errors"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"gorm.io/gorm"
	"time"
)

type BatchRepositoryImpl struct {
	db *gorm.DB
}

type BatchRecord struct {
	gorm.Model
	BatchName         string
	ServerName        string
	CronSetting       string
	InitialDate       time.Time
	TimeLimit         int
	EstimatedDuration int
}

func NewBatchRepository(db *gorm.DB) core.BatchRepository {
	return BatchRepositoryImpl{db: db}
}

func (BatchRecord) TableName() string {
	return "batch"
}

func (br BatchRepositoryImpl) FindById(id int) (*entity.Batch, error) {
	var b BatchRecord
	tx := br.db.Find(&b, id)
	if tx.RowsAffected == 0 {
		return nil, errors.New("no record")
	}

	batch := entity.Batch{
		Id:                 id,
		BatchName:          b.BatchName,
		ServerName:         b.ServerName,
		CronSetting:        b.CronSetting,
		TimeLimit:          b.TimeLimit,
		EsitimatedDuration: b.EstimatedDuration,
		StartDate:          b.InitialDate,
	}
	return &batch, nil
}
