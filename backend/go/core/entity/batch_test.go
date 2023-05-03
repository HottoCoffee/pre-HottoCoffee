package entity_test

import (
	"github.com/HottoCoffee/HottoCoffee/util"
	"reflect"
	"testing"
	"time"

	"github.com/HottoCoffee/HottoCoffee/core/entity"
)

func TestNewBatch(t *testing.T) {
	type args struct {
		id                 int
		batchName          string
		serverName         string
		cronSetting        string
		timeLimit          int
		estimationDuration int
		startDate          time.Time
		endDate            *time.Time
	}
	tests := []struct {
		name       string
		args       args
		want       *entity.Batch
		wantErr    bool
		wantErrMsg *string
	}{
		{
			"normal scenario without endDate",
			args{1, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			&entity.Batch{1, "batch", "server", newCronSetting("* * * * *"), 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			false,
			nil,
		},
		{
			"normal scenario with endDate",
			args{1, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), &[]time.Time{time.Date(2023, 1, 2, 0, 0, 0, 0, util.JST)}[0]},
			&entity.Batch{1, "batch", "server", newCronSetting("* * * * *"), 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), &[]time.Time{time.Date(2023, 1, 2, 0, 0, 0, 0, util.JST)}[0]},
			false,
			nil,
		},
		{
			"error scenario with invalid id",
			args{0, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			nil,
			true,
			&[]string{"ID should be more than 0. Given: 0"}[0],
		},
		{
			"error scenario with empty batch name",
			args{1, "", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			nil,
			true,
			&[]string{"batch name should not be empty"}[0],
		},
		{
			"error scenario with empty server name",
			args{1, "batch", "", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			nil,
			true,
			&[]string{"server name should not be empty"}[0],
		},
		{
			"error scenario with invalid cron setting",
			args{1, "batch", "server", "* * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			nil,
			true,
			&[]string{"malformed schedule setting * * * *"}[0],
		},
		{
			"error scenario with invalid time limit",
			args{1, "batch", "server", "* * * * *", 0, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			nil,
			true,
			&[]string{"time limit should be equal or more than 1. Given: 0"}[0],
		},
		{
			"error scenario with invalid estimation duration",
			args{1, "batch", "server", "* * * * *", 2, 3, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil},
			nil,
			true,
			&[]string{"estimation duration should be equal or more than 0 and less than time limit"}[0],
		},
		{
			"error scenario with invalid end date",
			args{1, "batch", "server", "* * * * *", 1, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), &[]time.Time{time.Date(2022, 12, 31, 0, 0, 0, 0, util.JST)}[0]},
			nil,
			true,
			&[]string{"end date should be after start date if exists"}[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewBatch(tt.args.id, tt.args.batchName, tt.args.serverName, tt.args.cronSetting, tt.args.timeLimit, tt.args.estimationDuration, tt.args.startDate, tt.args.endDate)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewBatch() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if *tt.wantErrMsg != err.Error() {
					t.Errorf("NewBatch() error message = %v, wantMsg %v", err.Error(), *tt.wantErrMsg)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newCronSetting(v string) entity.CronSetting {
	s, _ := entity.NewCronSetting(v)
	return *s
}
