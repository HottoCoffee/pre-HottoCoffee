package usecase

import (
	"github.com/HottoCoffee/HottoCoffee/util"
	"reflect"
	"testing"
	"time"
)

func Test_validateBatchInput(t *testing.T) {
	type args struct {
		input BatchInput
	}
	tests := []struct {
		name       string
		args       args
		want       *BatchInput
		wantErr    bool
		wantErrMsg string
	}{
		{
			"validation success",
			args{BatchInput{BatchName: "batch", ServerName: "server", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), TimeLimit: 1, CronSetting: "* * * * *"}},
			&BatchInput{BatchName: "batch", ServerName: "server", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), TimeLimit: 1, CronSetting: "* * * * *"},
			false,
			"",
		}, {
			"validation error",
			args{BatchInput{}},
			nil,
			true,
			"batch name should not be empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateBatchInput(tt.args.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("validateBatchInput() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != tt.wantErrMsg {
					t.Errorf("validateBatchInput() error = %v, wantErrMsg %v", err, tt.wantErrMsg)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateBatchInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
