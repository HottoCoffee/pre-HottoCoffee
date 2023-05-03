package usecase

import (
	"errors"
	mock_core "github.com/HottoCoffee/HottoCoffee/.mock/core"
	mock_usecase "github.com/HottoCoffee/HottoCoffee/.mock/usecase"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestCreateBatchUsecase_Execute(t *testing.T) {
	type args struct {
		input BatchInput
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary)
	}{
		{
			"successfully created",
			args{BatchInput{BatchName: "batch", ServerName: "server", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), TimeLimit: 1, CronSetting: "* * * * *"}},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				br.EXPECT().Create("batch", "server", "* * * * *", 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)).
					Return(entity.NewBatch(1, "batch", "server", "* * * * *", 1, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil))
				batch, _ := entity.NewBatch(1, "batch", "server", "* * * * *", 1, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), nil)
				bob.EXPECT().SendBatchResponse(*batch)
			},
		}, {
			"validation error",
			args{BatchInput{}},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				bob.EXPECT().SendInvalidRequestResponse("batch name should not be empty").Times(1)
				br.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
		}, {
			"insertion error",
			args{BatchInput{BatchName: "batch", ServerName: "server", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), TimeLimit: 1, CronSetting: "* * * * *"}},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				br.EXPECT().Create("batch", "server", "* * * * *", 1, time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)).
					Return(nil, errors.New("sample error"))
				bob.EXPECT().SendInternalServerErrorResponse().Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			br := mock_core.NewMockBatchRepository(ctrl)
			bob := mock_usecase.NewMockBatchOutputBoundary(ctrl)
			tt.prepareMockFn(br, bob)

			cbu := NewCreateBatchUsecase(br, bob)
			cbu.Execute(tt.args.input)
		})
	}
}
