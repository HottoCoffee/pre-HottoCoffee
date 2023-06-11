package entity

import (
	"github.com/HottoCoffee/HottoCoffee/util"
	"reflect"
	"testing"
	"time"
)

func TestNewBatchExecutionHistories(t *testing.T) {
	batch, _ := NewBatch(1, "batch", "server", "0 0 * * *", 60, 0,
		time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
		&[]time.Time{time.Date(2024, 1, 1, 0, 0, 0, 0, util.JST)}[0],
	)

	type args struct {
		histories []History
	}
	tests := []struct {
		name string
		args args
		want BatchExecutionHistories
	}{
		{
			"normal scenario",
			args{[]History{
				{1, "success", time.Date(2023, 1, 1, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
				{2, "failed", time.Date(2023, 1, 2, 0, 0, 1, 1, util.JST), time.Date(2023, 1, 2, 1, 1, 0, 0, util.JST)},
				{3, "success", time.Date(2023, 1, 4, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
				{4, "success", time.Date(2023, 1, 3, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
			}},
			BatchExecutionHistories{Batch: *batch, Histories: []History{
				{1, "success", time.Date(2023, 1, 1, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
				{2, "failed", time.Date(2023, 1, 2, 0, 0, 1, 1, util.JST), time.Date(2023, 1, 2, 1, 1, 0, 0, util.JST)},
				{4, "success", time.Date(2023, 1, 3, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
				{3, "success", time.Date(2023, 1, 4, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
			}},
		},
		{
			"if execution results are duplicated for the single batch execution, select the latest one",
			args{[]History{
				{1, "success", time.Date(2023, 1, 1, 0, 0, 2, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
				{2, "failed", time.Date(2023, 1, 1, 23, 59, 0, 0, util.JST), time.Date(2023, 1, 2, 0, 59, 0, 0, util.JST)},
			}},
			BatchExecutionHistories{Batch: *batch, Histories: []History{
				{2, "failed", time.Date(2023, 1, 1, 23, 59, 0, 0, util.JST), time.Date(2023, 1, 2, 0, 59, 0, 0, util.JST)}}},
		},
		{
			"omit execution results that are executed before batch start or after batch end",
			args{[]History{
				{1, "success", time.Date(2022, 12, 31, 23, 59, 59, 0, util.JST), time.Date(2023, 1, 1, 0, 50, 0, 0, util.JST)},
				{2, "failed", time.Date(2024, 1, 1, 0, 0, 1, 0, util.JST), time.Date(2024, 1, 1, 0, 59, 0, 0, util.JST)},
				{3, "failed", time.Date(2024, 1, 2, 0, 0, 1, 0, util.JST), time.Date(2024, 1, 2, 0, 59, 0, 0, util.JST)},
			}},
			BatchExecutionHistories{Batch: *batch, Histories: []History{
				{2, "failed", time.Date(2024, 1, 1, 0, 0, 1, 0, util.JST), time.Date(2024, 1, 1, 0, 59, 0, 0, util.JST)}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBatchExecutionHistories(*batch, tt.args.histories); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBatchExecutionHistories() = %v, want %v", got, tt.want)
			}
		})
	}
}
