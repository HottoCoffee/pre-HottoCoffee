package infrastructure

import (
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/HottoCoffee/HottoCoffee/util"
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

func TestHistoryRepositoryImpl_FindByHistoryIdAndBatchId(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	type fields struct {
		db gorm.DB
	}
	type args struct {
		historyId int
		batchId   int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		inputDate  func()
		want       *entity.BatchExecutionHistory
		wantErr    bool
		wantErrMsg *string
	}{
		{
			"normal scenario",
			fields{*db},
			args{historyId: 1, batchId: 1},
			func() {
				db.Exec(`insert into batch values (1, 'batch', 'server', '0 1 * * *', '2023-01-01', 10, 0, '2022-01-01', '2022-01-01', null)`)
				db.Exec(`insert into history values (1, 1, 'success', '2023-01-01', '2023-01-01', '2023-01-01', '2023-01-01', null)`)
			},
			&entity.BatchExecutionHistory{
				Batch: entity.Batch{
					Id:                1,
					BatchName:         "batch",
					ServerName:        "server",
					CronSetting:       newCronSetting("0 1 * * *"),
					TimeLimit:         10,
					EstimatedDuration: 0,
					StartDate:         time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
					EndDate:           nil,
				},
				History: entity.History{
					Id:              1,
					ExecutionResult: entity.Success,
					StartDatetime:   time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
					FinishDatetime:  time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
				},
			},
			true,
			nil,
		}, {
			"normal scenario with deleted batch",
			fields{*db},
			args{historyId: 1, batchId: 1},
			func() {
				db.Exec(`insert into batch values (1, 'batch', 'server', '0 1 * * *', '2023-01-01', 10, 0, '2023-01-01', '2023-01-01', '2023-02-01')`)
				db.Exec(`insert into history values (1, 1, 'success', '2023-01-01', '2023-01-01', '2023-01-01', '2023-01-01', null)`)
			},
			&entity.BatchExecutionHistory{
				Batch: entity.Batch{
					Id:                1,
					BatchName:         "batch",
					ServerName:        "server",
					CronSetting:       newCronSetting("0 1 * * *"),
					TimeLimit:         10,
					EstimatedDuration: 0,
					StartDate:         time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
					EndDate:           &[]time.Time{time.Date(2023, 2, 1, 0, 0, 0, 0, util.JST)}[0],
				},
				History: entity.History{
					Id:              1,
					ExecutionResult: entity.Success,
					StartDatetime:   time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
					FinishDatetime:  time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
				},
			},
			false,
			nil,
		}, {
			"error scenario with domain violation",
			fields{*db},
			args{historyId: 1, batchId: 1},
			func() {
				db.Exec(`insert into batch values (1, 'batch', 'server', '0 1 * * *', '2023-01-01', 10, 0, '2023-01-01', '2023-01-01', '2023-02-01')`)
				db.Exec(`insert into history values (1, 1, 'success', '2022-01-01', '2022-01-01', '2022-01-01', '2022-01-01', null)`)
			},
			nil, true,
			&[]string{"history should be started after batch's start date, batch id: 1, history id: 1, batch start datetime: 2023-01-01 00:00:00 +0900 JST, history start datetime: 2022-01-01 00:00:00 +0900 JST"}[0],
		},
	}
	for _, tt := range tests {
		truncate()
		tt.inputDate()
		t.Run(tt.name, func(t *testing.T) {
			hr := HistoryRepositoryImpl{
				db: tt.fields.db,
			}
			got, err := hr.FindByHistoryIdAndBatchId(tt.args.historyId, tt.args.batchId)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("FindByHistoryIdAndBatchId() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != *tt.wantErrMsg {
					t.Errorf("FindByHistoryIdAndBatchId() error message = %v, wantErrMsg %v", err.Error(), *tt.wantErrMsg)
					return
				}
				return
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(*got, "Batch.EndDate"), cmpopts.IgnoreUnexported(got.Batch.CronSetting)); len(diff) != 0 {
				t.Errorf("FindByHistoryIdAndBatchId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistoryRepositoryImpl_FindByBatchId(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	type fields struct {
		db gorm.DB
	}
	type args struct {
		batchId int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		inputDate  func()
		test       func(histories entity.BatchExecutionHistories) bool
		wantErr    bool
		wantErrMsg *string
	}{
		{
			"normal scenario",
			fields{*db},
			args{1},
			func() {
				db.Exec(`insert into batch values (1, 'batch', 'server', '0 0 * * *', '2023-01-01', 10, 0, '2022-01-01', '2022-01-01', null)`)
				db.Exec(`insert into history values (1, 1, 'success', '2023-01-01', '2023-01-01 00:10:00', '2023-01-01', '2023-01-01', null)`)
				db.Exec(`insert into history values (2, 1, 'success', '2023-01-02', '2023-01-02 00:10:00', '2023-01-02', '2023-01-02', null)`)
			},
			func(histories entity.BatchExecutionHistories) bool {
				return histories.Batch.Id == 1 &&
					len(histories.Histories) == 2
			},
			false,
			nil,
		},
	}
	for _, tt := range tests {
		truncate()
		tt.inputDate()
		t.Run(tt.name, func(t *testing.T) {
			hr := HistoryRepositoryImpl{
				db: tt.fields.db,
			}
			got, err := hr.FindByBatchId(tt.args.batchId)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("FindByBatchId() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != *tt.wantErrMsg {
					t.Errorf("FindByBatchId() error message = %v, wantErrMsg %v", err.Error(), *tt.wantErrMsg)
					return
				}
				return
			}
			if !tt.test(*got) {
				t.Errorf("FindByBatchId()")
			}
		})
	}
}

func newCronSetting(value string) entity.CronSetting {
	setting, _ := entity.NewCronSetting(value)
	return *setting
}
