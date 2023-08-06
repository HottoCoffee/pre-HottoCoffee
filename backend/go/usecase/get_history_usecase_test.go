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

func TestGetHistoryUsecase_Execute(t *testing.T) {
	type args struct {
		batchIdString   string
		historyIdString string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary)
	}{
		{
			"normal scenario",
			args{batchIdString: "1", historyIdString: "1"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				batch, _ := entity.NewBatch(1, "batch", "server", "0 10 * * *", 10, 0, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), nil)
				history, _ := entity.NewHistory(1, "success", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), time.Date(2023, 1, 1, 0, 1, 0, 0, util.JST))
				executionHistory, _ := entity.NewBatchExecutionHistory(*batch, *history)

				hr.EXPECT().FindByHistoryIdAndBatchId(1, 1).Return(executionHistory, nil)

				hob.EXPECT().SendHistoryResponse(*executionHistory)
			},
		},
		{
			"error scenario with invalid batchIdString",
			args{batchIdString: "a", historyIdString: "1"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				hob.EXPECT().SendNotFoundResponse().Times(1)
			},
		},
		{
			"error scenario with invalid historyIdString",
			args{batchIdString: "1", historyIdString: "a"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				hob.EXPECT().SendNotFoundResponse().Times(1)
			},
		},
		{
			"error scenario with domain violated history",
			args{batchIdString: "1", historyIdString: "1"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				hr.EXPECT().FindByHistoryIdAndBatchId(1, 1).Return(nil, entity.NewDomainRuleViolationError("sample"))
				hob.EXPECT().SendInternalServerErrorResponse().Times(1)
			},
		},
		{
			"error scenario with not found history",
			args{batchIdString: "1", historyIdString: "1"},
			func(hr *mock_core.MockHistoryRepository, hob *mock_usecase.MockHistoryOutputBoundary) {
				hr.EXPECT().FindByHistoryIdAndBatchId(1, 1).Return(nil, errors.New("not found"))
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

			ghu := NewGetHistoryUsecase(hr, hob)
			ghu.Execute(tt.args.batchIdString, tt.args.historyIdString)
		})
	}
}
