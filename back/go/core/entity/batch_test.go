package entity_test

import (
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
		name    string
		args    args
		want    *entity.Batch
		wantErr bool
	}{
		{
			"normal scenario without endDate",
			args{1, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			&entity.Batch{1, "batch", "server", newCronSetting("* * * * *"), 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			false,
		},
		{
			"normal scenario with endDate",
			args{1, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), &[]time.Time{time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)}[0]},
			&entity.Batch{1, "batch", "server", newCronSetting("* * * * *"), 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), &[]time.Time{time.Date(2023, 1, 2, 0, 0, 0, 0, time.Local)}[0]},
			false,
		},
		{
			"error scenario with invalid id",
			args{0, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			nil,
			true,
		},
		{
			"error scenario with empty batch name",
			args{1, "", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			nil,
			true,
		},
		{
			"error scenario with empty server name",
			args{1, "batch", "", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			nil,
			true,
		},
		{
			"error scenario with invalid cron setting",
			args{1, "batch", "server", "* * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			nil,
			true,
		},
		{
			"error scenario with invalid time limit",
			args{1, "batch", "server", "* * * * *", 0, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			nil,
			true,
		},
		{
			"error scenario with invalid estimation duration",
			args{1, "batch", "server", "* * * * *", 2, 3, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil},
			nil,
			true,
		},
		{
			"error scenario with invalid end date",
			args{1, "batch", "server", "* * * * *", 2, 3, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), &[]time.Time{time.Date(2022, 12, 31, 0, 0, 0, 0, time.Local)}[0]},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := entity.NewBatch(tt.args.id, tt.args.batchName, tt.args.serverName, tt.args.cronSetting, tt.args.timeLimit, tt.args.estimationDuration, tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
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
