package entity

import (
	"errors"
	"testing"
)

func TestIsDomainRuleViolationError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"domain violation error",
			args{NewDomainRuleViolationError("sample")},
			true,
		}, {
			"normal error",
			args{errors.New("sample")},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDomainRuleViolationError(tt.args.err); got != tt.want {
				t.Errorf("IsDomainRuleViolationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
