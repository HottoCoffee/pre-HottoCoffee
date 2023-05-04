package infrastructure

import (
	"errors"
	"fmt"
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
	for _, br := range brs {
		batch, err := mapRecordToBatch(br)
		if err != nil {
			return nil, errors.New("broken DB records")
		}
		bs = append(bs, *batch)
	}
	return bs, nil
}

func (br BatchRepositoryImpl) FindFilteredBy(query string) ([]entity.Batch, error) {
	var brs []BatchRecord
	br.db.Where("batch_name LIKE ?", "%"+query+"%").
		Or("server_name LIKE ?", "%"+query+"%").
		Or("cron_setting LIKE ?", "%"+query+"%").
		Find(&brs)

	var bs []entity.Batch
	for _, br := range brs {
		batch, err := mapRecordToBatch(br)
		if err != nil {
			return nil, errors.New("broken DB records")
		}
		bs = append(bs, *batch)
	}
	return bs, nil
}

func (br BatchRepositoryImpl) Create(batchName string, serverName string, cronSetting string, timeLimit int, startDate time.Time) (*entity.Batch, error) {
	batchRecord := BatchRecord{BatchName: batchName, ServerName: serverName, CronSetting: cronSetting, InitialDate: startDate, TimeLimit: timeLimit}
	tx := br.db.Create(&batchRecord)
	if tx.RowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("failed to insert batch record: %v", batchRecord))
	}

	return mapRecordToBatch(batchRecord)
}

func (br BatchRepositoryImpl) Save(batch entity.Batch) error {
	record := BatchRecord{BatchName: batch.BatchName, ServerName: batch.ServerName, CronSetting: batch.CronSetting.ToString(), InitialDate: batch.StartDate, TimeLimit: batch.TimeLimit}
	record.ID = uint(batch.Id)
	tx := br.db.Omit("created_at").Save(&record)
	if tx.RowsAffected != 1 {
		return errors.New(fmt.Sprintf("failed to update batch record: %v", record))
	}
	return nil
}

func mapRecordToBatch(r BatchRecord) (*entity.Batch, error) {
	var endDate *time.Time = nil
	if r.DeletedAt.Valid {
		endDate = &r.DeletedAt.Time
	}
	return entity.NewBatch(int(r.ID), r.BatchName, r.ServerName, r.CronSetting, r.TimeLimit, r.EstimatedDuration, r.InitialDate, endDate)
}
