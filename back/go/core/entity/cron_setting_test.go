package entity

import (
	"reflect"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestNewCronSetting(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    *CronSetting
		wantErr bool
	}{
		{"invalid cron setting", args{"* * * *"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCronSetting(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCronSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCronSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCronSetting_ListSchedules(t *testing.T) {
	type fields struct {
		value    string
		schedule cron.Schedule
	}
	type args struct {
		begin        time.Time
		endInclusive time.Time
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantSchedules []time.Time
	}{
		{
			"",
			fields{"*/30 * * * *", *newSchedule("*/30 * * * *")},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local)},
			[]time.Time{
				time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
				time.Date(2023, 1, 1, 0, 30, 0, 0, time.Local),
				time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
			},
		},
		{
			"same time",
			fields{"*/30 * * * *", *newSchedule("*/30 * * * *")},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)},
			[]time.Time{time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)},
		},
		{
			"empty list",
			fields{"0 2 * * *", *newSchedule("0 2 * * *")},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)},
			[]time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CronSetting{
				value:    tt.fields.value,
				schedule: tt.fields.schedule,
			}
			gotSchedules := cs.ListSchedules(tt.args.begin, tt.args.endInclusive)

			if len(tt.wantSchedules) == 0 {
				if !(len(gotSchedules) == 0) {
					t.Errorf("CronSetting.ListSchedules() = %v, want %v", gotSchedules, tt.wantSchedules)
				}
			} else {
				if !reflect.DeepEqual(gotSchedules, tt.wantSchedules) {
					t.Errorf("CronSetting.ListSchedules() = %v, want %v", gotSchedules, tt.wantSchedules)
				}
			}
		})
	}
}

func newSchedule(v string) *cron.Schedule {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, _ := p.Parse(v)
	return &s
}
