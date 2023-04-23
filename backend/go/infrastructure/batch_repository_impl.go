package infrastructure

import (
	"errors"
	"time"

	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"gorm.io/gorm"
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

	var da *time.Time
	if b.DeletedAt.Valid {
		da = &b.DeletedAt.Time
	}

	batch, err := entity.NewBatch(id, b.BatchName, b.ServerName, b.CronSetting, b.TimeLimit, b.EstimatedDuration, b.InitialDate, da)
	if err != nil {
		return nil, errors.New("broken DB record")
	}
	return batch, nil
}

func (br BatchRepositoryImpl) FindAll() ([]entity.Batch, error) {
	var brs []BatchRecord
	br.db.Find(&brs)

	var bs []entity.Batch
	for i := range brs {
		batch, err := entity.NewBatch(int(brs[i].ID), brs[i].BatchName, brs[i].ServerName, brs[i].CronSetting, brs[i].TimeLimit, brs[i].EstimatedDuration, brs[i].InitialDate, nil)
		if err != nil {
			return nil, errors.New("broken DB records")
		}
		bs = append(bs, *batch)
	}
	return bs, nil
}

func (br BatchRepositoryImpl) FindFilteredBy(query string) ([]entity.Batch, error) {
	return nil, errors.New("not impletemented")
}
