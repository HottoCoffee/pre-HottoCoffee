package entity

import (
	"github.com/HottoCoffee/HottoCoffee/util"
	"reflect"
	"testing"
	"time"
)

func TestNewHistory(t *testing.T) {
	type args struct {
		id                    int
		executionResultString string
		startDatetime         time.Time
		finishDatetime        time.Time
	}
	tests := []struct {
		name       string
		args       args
		want       *History
		wantErr    bool
		wantErrMsg *string
	}{
		{
			name: "normal scenario with success status",
			args: args{1, "success", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST)},
			want: &History{
				Id:              1,
				ExecutionResult: Success,
				StartDatetime:   time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
				FinishDatetime:  time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST),
			},
			wantErr:    false,
			wantErrMsg: nil,
		},
		{
			name: "normal scenario with failed status",
			args: args{1, "failed", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST)},
			want: &History{
				Id:              1,
				ExecutionResult: Failed,
				StartDatetime:   time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST),
				FinishDatetime:  time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST),
			},
			wantErr:    false,
			wantErrMsg: nil,
		},
		{
			name:       "error scenario with invalid id",
			args:       args{0, "success", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST)},
			want:       nil,
			wantErr:    true,
			wantErrMsg: &[]string{"history id should be equal or greater than 0. given: 0"}[0],
		},
		{
			name:       "error scenario with invalid status",
			args:       args{1, "successful", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST)},
			want:       nil,
			wantErr:    true,
			wantErrMsg: &[]string{"execution result should be success or failed. given: successful"}[0],
		},
		{
			name:       "error scenario with invalid date",
			args:       args{1, "success", time.Date(2023, 1, 1, 0, 2, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST)},
			want:       nil,
			wantErr:    true,
			wantErrMsg: &[]string{"history finish datetime should be after start datetime. start datetime: 2023-01-01 00:02:00 +0900 JST, finish datetime: 2023-01-01 00:01:00 +0900 JST"}[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHistory(tt.args.id, tt.args.executionResultString, tt.args.startDatetime, tt.args.finishDatetime)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewHistory() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != *tt.wantErrMsg {
					t.Errorf("NewHistory() error message = %v, wantErrMsg %v", err.Error(), *tt.wantErrMsg)
					return
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
