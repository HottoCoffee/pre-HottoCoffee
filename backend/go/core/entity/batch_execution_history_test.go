package entity

import (
	"github.com/HottoCoffee/HottoCoffee/util"
	"reflect"
	"testing"
	"time"
)

func TestNewBatchExecutionHistory(t *testing.T) {
	type args struct {
		batch   Batch
		history History
	}
	tests := []struct {
		name       string
		args       args
		want       *BatchExecutionHistory
		wantErr    bool
		wantErrMsg *string
	}{
		{
			"normal scenario",
			args{batch: batch, history: newHistory(time.Date(2023, 4, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 4, 1, 0, 1, 23, 45, util.JST))},
			&BatchExecutionHistory{batch, newHistory(time.Date(2023, 4, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 4, 1, 0, 1, 23, 45, util.JST))},
			false,
			nil,
		},
		{
			"normal scenario with finished after time limit",
			args{batch: batch, history: newHistory(time.Date(2023, 4, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 4, 1, 1, 0, 0, 0, util.JST))},
			&BatchExecutionHistory{batch, newHistory(time.Date(2023, 4, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 4, 1, 1, 0, 0, 0, util.JST))},
			false,
			nil,
		}, {
			"error scenario with started before batch start date",
			args{batch: batch, history: newHistory(time.Date(2022, 12, 31, 23, 59, 59, 99, util.JST), time.Date(2023, 1, 1, 0, 1, 23, 45, util.JST))},
			nil,
			true,
			&[]string{"history should be started after batch's start date, batch id: 1, history id: 1, batch start datetime: 2023-01-01 00:00:00 +0900 JST, history start datetime: 2022-12-31 23:59:59.000000099 +0900 JST"}[0],
		}, {
			"error scenario with started after batch end date",
			args{batch: batch, history: newHistory(time.Date(2024, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2024, 1, 1, 0, 1, 23, 45, util.JST))},
			nil,
			true,
			&[]string{"history should not be started after batch's end date, batch id: 1, history id: 1, batch start datetime: 2023-01-01 00:00:00 +0900 JST, history start datetime: 2024-01-01 00:00:00 +0900 JST"}[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBatchExecutionHistory(tt.args.batch, tt.args.history)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewBatchExecutionHistory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != *tt.wantErrMsg {
					t.Errorf("NewBatchExecutionHistory() error message = %v, wantErrMsg %v", err.Error(), *tt.wantErrMsg)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBatchExecutionHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var batch = Batch{
	Id:                1,
	BatchName:         "batch",
	ServerName:        "server",
	CronSetting:       newPerHourCronSetting(),
	TimeLimit:         10,
	EstimatedDuration: 0,
	StartDate:         time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
	EndDate:           &[]time.Time{time.Date(2023, 12, 31, 0, 0, 0, 0, util.JST)}[0],
}

func newPerHourCronSetting() CronSetting {
	setting, _ := NewCronSetting("0 1 * * *")
	return *setting
}

func newHistory(startDatetime time.Time, finishDatetime time.Time) History {
	return History{
		Id:              1,
		ExecutionResult: "success",
		StartDatetime:   startDatetime,
		FinishDatetime:  finishDatetime,
	}
}
