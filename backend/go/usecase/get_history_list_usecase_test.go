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

func TestGetHistoryListUsecase_Execute(t *testing.T) {
	type args struct {
		batchIdString string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary)
	}{
		{
			"normal scenario",
			args{"1"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				batch := newBatch(1, "batch", "server", "* * * * *", 2, 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil)
				histories := []entity.History{
					{1, "success", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 0, 10, 0, util.JST)},
					{2, "success", time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 10, 0, util.JST)},
				}

				executionHistories := entity.NewBatchExecutionHistories(batch, histories)
				hr.EXPECT().FindByBatchId(1).Return(&executionHistories, nil).Times(1)
				hob.EXPECT().SendHistoryListResponse(executionHistories).Times(1)
			},
		},
		{
			"batchId is not integer",
			args{"a"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				hr.EXPECT().FindByBatchId(gomock.Any()).Times(0)
				hob.EXPECT().SendNotFoundResponse().Times(1)
			},
		},
		{
			"target batch is not found",
			args{"1"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				hr.EXPECT().FindByBatchId(1).Return(nil, errors.New("not found")).Times(1)
				hob.EXPECT().SendNotFoundResponse().Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			hr := mock_core.NewMockHistoryRepository(ctrl)
			hob := mock_usecase.NewMockHistoryOutputBoundary(ctrl)
			tt.prepareMockFn(hr, hob)

			ghlu := NewGetHistoryListUsecase(hr, hob)
			ghlu.Execute(tt.args.batchIdString)
		})
	}
}
