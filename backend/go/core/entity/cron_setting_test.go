package entity

import (
	"github.com/HottoCoffee/HottoCoffee/util"
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
		name       string
		args       args
		want       *CronSetting
		wantErr    bool
		wantErrMsg *string
	}{
		{
			"normal scenario",
			args{"* * * * *"},
			&CronSetting{"* * * * *", *newSchedule("* * * * *")},
			false,
			nil,
		},
		{
			"normal scenario with invalid cron setting",
			args{"* * * *"},
			nil,
			true,
			&[]string{"malformed schedule setting * * * *"}[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCronSetting(tt.args.v)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewCronSetting() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if *tt.wantErrMsg != err.Error() {
					t.Errorf("NewCronSetting() error message = %v, wantMsg %v", err.Error(), *tt.wantErrMsg)
					return
				}
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
			"normal scenario",
			fields{"*/30 * * * *", *newSchedule("*/30 * * * *")},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 1, 0, 0, 0, util.JST)},
			[]time.Time{
				time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
				time.Date(2023, 1, 1, 0, 30, 0, 0, util.JST),
				time.Date(2023, 1, 1, 1, 0, 0, 0, util.JST),
			},
		},
		{
			"begin time and end time is the same value",
			fields{"*/30 * * * *", *newSchedule("*/30 * * * *")},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST)},
			[]time.Time{time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST)},
		},
		{
			"there are no scheduled batch executions and return empty list",
			fields{"0 2 * * *", *newSchedule("0 2 * * *")},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 1, 0, 0, 0, util.JST)},
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
				if len(gotSchedules) != 0 {
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

func TestCronSetting_Prev(t *testing.T) {
	type fields struct {
		value string
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{"every minutes",
			fields{"* * * * *"},
			args{time.Date(2023, 1, 2, 3, 4, 5, 6, util.JST)},
			time.Date(2023, 1, 2, 3, 4, 0, 0, util.JST),
		},
		{"every hour",
			fields{"0 * * * *"},
			args{time.Date(2023, 1, 2, 3, 4, 5, 6, util.JST)},
			time.Date(2023, 1, 2, 3, 0, 0, 0, util.JST),
		},
		{"every day",
			fields{"0 0 * * *"},
			args{time.Date(2023, 1, 2, 3, 4, 5, 6, util.JST)},
			time.Date(2023, 1, 2, 0, 0, 0, 0, util.JST),
		},
		{"every month",
			fields{"0 0 1 * *"},
			args{time.Date(2023, 1, 2, 3, 4, 5, 6, util.JST)},
			time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
		},
		{"every Sunday",
			fields{"0 0 * * 0"},
			args{time.Date(2023, 1, 2, 3, 4, 5, 6, util.JST)}, // Monday
			time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
		},
		{"over year",
			fields{"* * * * *"},
			args{time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST)},
			time.Date(2022, 12, 31, 23, 59, 0, 0, util.JST),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := newCronSetting(tt.fields.value)
			if got := cs.Prev(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newCronSetting(v string) CronSetting {
	cs, _ := NewCronSetting(v)
	return *cs
}

func newSchedule(v string) *cron.Schedule {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, _ := p.Parse(v)
	return &s
}
