package usecase_test

import (
	"errors"
	"testing"
	"time"

	mock_core "github.com/HottoCoffee/HottoCoffee/.mock/core"
	mock_usecase "github.com/HottoCoffee/HottoCoffee/.mock/usecase"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/HottoCoffee/HottoCoffee/usecase"
	"github.com/golang/mock/gomock"
)

func TestGetBatchUsecase_Execute(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary)
	}{
		{
			"normal scenario",
			args{"1"},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				batch, _ := entity.NewBatch(1, "batch", "server", "* * * * *", 2, 1, time.Now(), nil)

				br.EXPECT().
					FindById(1).
					Return(batch, nil)

				bob.EXPECT().
					SendBatchResponse(*batch)
			},
		},
		{
			"input value is invalid and return not found response",
			args{"string"},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				br.EXPECT().
					FindById(gomock.Any()).
					Times(0)

				bob.EXPECT().
					SendNotFoundResponse()
			},
		},
		{
			"matching data does not exist in db and return not found response",
			args{"1"},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				br.EXPECT().
					FindById(1).
					Return(nil, errors.New("no recode"))

				bob.EXPECT().
					SendNotFoundResponse()
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

			gbu := usecase.NewGetBatchUsecase(br, bob)
			gbu.Execute(tt.args.input)
		})
	}
}
