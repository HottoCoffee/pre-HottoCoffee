package usecase

import (
	"github.com/HottoCoffee/HottoCoffee/util"
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
		wantErr    bool
		wantErrMsg string
	}{
		{
			"validation success",
			args{BatchInput{BatchName: "batch", ServerName: "server", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), TimeLimit: 1, CronSetting: "* * * * *"}},
			false,
			"",
		}, {
			"validation error",
			args{BatchInput{}},
			true,
			"batch name should not be empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBatchInput(tt.args.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("validateBatchInput() error = %v, wantErr %v", err, tt.wantErr)
				} else if err.Error() != tt.wantErrMsg {
					t.Errorf("validateBatchInput() error = %v, wantErrMsg %v", err, tt.wantErrMsg)
				}
				return
			}
			if err != nil {
				t.Errorf("validateBatchInput() err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
