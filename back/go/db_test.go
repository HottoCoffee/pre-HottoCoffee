// TODO: This test will be removed after merging https://github.com/HottoCoffee/HottoCoffee/pull/10
// because this is just for checking DB spec.

package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

var dialector = mysql.Open("root:root@tcp(127.0.0.1)/hottocoffee?parseTime=true&loc=Asia%2FTokyo")
var db, _ = gorm.Open(dialector, &gorm.Config{})

func truncate() {
	db.Exec("set foreign_key_checks = 0")
	db.Exec("truncate table batch")
	db.Exec("truncate table history")
	db.Exec("set foreign_key_checks = 1")
}

func TestInsertOnBatch(t *testing.T) {
	// given
	if testing.Short() {
		t.Skip()
	}
	truncate()

	batch := Batch{BatchName: "hoge", ServerName: "fuga", CronSetting: "piyo", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), TimeLimit: 100, EstimatedDuration: 50}

	// when
	db.Create(&batch)

	// then:
	initialTime := time.Time{}
	if batch.ID == 0 || batch.CreatedAt == initialTime || batch.UpdatedAt == initialTime {
		t.Error("inserted batch struct was not updated by GORM")
	}

	// select
	var got Batch
	db.First(&got)
	if diff := cmp.Diff(got, batch, cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt", "DeletedAt")); len(diff) != 0 {
		t.Errorf("differs: (-got +want)\n%s", diff)
	}
}

func TestUpdateOnBatch(t *testing.T) {
	// given
	if testing.Short() {
		t.Skip()
	}
	truncate()

	batch := Batch{BatchName: "hoge", ServerName: "fuga", CronSetting: "piyo", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), TimeLimit: 100, EstimatedDuration: 50}
	db.Create(&batch)

	var originalRecord Batch
	db.First(&originalRecord)

	originalCreatedAt := originalRecord.CreatedAt
	originalUpdatedAt := originalRecord.UpdatedAt

	time.Sleep(time.Second) // wait 1 sec to confirm UpdatedAt is actually changed by update query

	// when
	originalRecord.BatchName = "foo"
	originalRecord.ServerName = "bar"
	originalRecord.CronSetting = "baz"
	db.Save(originalRecord)

	// then
	var got Batch
	db.First(&got)
	if diff := cmp.Diff(got, originalRecord, cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt", "DeletedAt")); len(diff) != 0 {
		t.Errorf("differs: (-got +want)\n%s", diff)
	}
	if got.CreatedAt != originalCreatedAt {
		t.Errorf("CreatedAt value should not be changed by update query\noriginal: %s\nupdated: %s", originalCreatedAt, got.CreatedAt)
	}
	if got.UpdatedAt == originalUpdatedAt {
		t.Error("UpdatedAt value should be changed by update query")
	}
}

func TestDeleteOnBatch(t *testing.T) {
	// given
	if testing.Short() {
		t.Skip()
	}
	truncate()

	batch := Batch{BatchName: "hoge", ServerName: "fuga", CronSetting: "piyo", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), TimeLimit: 100, EstimatedDuration: 50}
	db.Create(&batch)

	// when
	db.Delete(&batch)

	// then
	if !batch.DeletedAt.Valid {
		t.Error("DeletedAt value should be made valid by delete query")
	}

	var rows []Batch
	db.Find(&rows)
	if len(rows) != 0 {
		t.Error("Normal select query should not fetch logical-deleted records")
	}

	db.Unscoped().Find(&rows)
	if len(rows) == 0 {
		t.Error("Unscoped select query should fetch logical-deleted records")
	}
}

func TestInsertOnHistory(t *testing.T) {
	// given
	if testing.Short() {
		t.Skip()
	}
	truncate()

	db.Create(&Batch{InitialDate: time.Now()})
	history := History{BatchId: 1, Status: "success"}

	// when
	db.Create(&history)

	// then
	var got History
	db.First(&got)
	if diff := cmp.Diff(got, history, cmpopts.IgnoreFields(got, "CreatedAt", "UpdatedAt", "DeletedAt")); len(diff) != 0 {
		t.Errorf("differs: (-got +want)\n%s", diff)
	}
}

type Batch struct {
	gorm.Model
	BatchName         string
	ServerName        string
	CronSetting       string
	InitialDate       time.Time
	TimeLimit         int
	EstimatedDuration int
}

func (Batch) TableName() string {
	return "batch"
}

type History struct {
	gorm.Model
	BatchId int
	Status  string
}

func (History) TableName() string {
	return "history"
}
