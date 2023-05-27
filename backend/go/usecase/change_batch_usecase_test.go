package usecase

import (
	"errors"
	mock_core "github.com/HottoCoffee/HottoCoffee/.mock/core"
	mock_usecase "github.com/HottoCoffee/HottoCoffee/.mock/usecase"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/HottoCoffee/HottoCoffee/util"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestChangeBatchUsecase_Execute(t *testing.T) {
	type args struct {
		stringId string
		input    BatchInput
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary)
	}{
		{
			"successfully changed",
			args{
				stringId: "1",
				input: BatchInput{
					BatchName:   "batch2",
					ServerName:  "server2",
					InitialDate: time.Date(2024, 1, 1, 0, 0, 0, 0, util.JST),
					TimeLimit:   100,
					CronSetting: "2 * * * *",
				},
			},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				batch, _ := entity.NewBatch(1, "batch", "server", "* * * * *", 1, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil)
				br.EXPECT().FindById(1).Return(batch, nil)
				br.EXPECT().Save(gomock.Any()).Return(nil) // here, we cannot use batch as an arg. the arg does not match an actual one because cronSetting has pointer type.
				bob.EXPECT().SendBatchResponse(gomock.Any())
			},
		},
		{
			"stringId is not integer",
			args{
				stringId: "a",
				input: BatchInput{
					BatchName:   "batch2",
					ServerName:  "server2",
					InitialDate: time.Date(2024, 1, 1, 0, 0, 0, 0, util.JST),
					TimeLimit:   100,
					CronSetting: "2 * * * *",
				},
			},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				bob.EXPECT().SendNotFoundResponse().Times(1)
			},
		},
		{
			"no batch is found",
			args{
				stringId: "1",
				input: BatchInput{
					BatchName:   "batch2",
					ServerName:  "server2",
					InitialDate: time.Date(2024, 1, 1, 0, 0, 0, 0, util.JST),
					TimeLimit:   100,
					CronSetting: "2 * * * *",
				},
			},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				br.EXPECT().FindById(1).Return(nil, errors.New("not found"))
				bob.EXPECT().SendNotFoundResponse().Times(1)
			},
		},
		{
			"validation error",
			args{
				stringId: "1",
				input:    BatchInput{},
			},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				batch, _ := entity.NewBatch(1, "batch", "server", "* * * * *", 1, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil)
				br.EXPECT().FindById(1).Return(batch, nil)
				bob.EXPECT().SendInvalidRequestResponse("batch name should not be empty")
			},
		},
		{
			"update error",
			args{
				stringId: "1",
				input: BatchInput{
					BatchName:   "batch2",
					ServerName:  "server2",
					InitialDate: time.Date(2024, 1, 1, 0, 0, 0, 0, util.JST),
					TimeLimit:   100,
					CronSetting: "2 * * * *",
				},
			},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				batch, _ := entity.NewBatch(1, "batch", "server", "* * * * *", 1, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil)
				br.EXPECT().FindById(1).Return(batch, nil)
				br.EXPECT().Save(gomock.Any()).Return(errors.New("update error"))
				bob.EXPECT().SendInternalServerErrorResponse()
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

			cbu := NewChangeBatchUsecase(br, bob)
			cbu.Execute(tt.args.stringId, tt.args.input)
		})
	}
}
