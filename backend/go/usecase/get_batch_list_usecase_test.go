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

func TestGetBatchListUsecase_Execute(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary)
	}{
		{
			"normal scenario without query",
			args{""},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				bs := []entity.Batch{
					newBatch(1, "batch", "server", "* * * * *", 2, 1, time.Now(), nil),
					newBatch(2, "batch2", "server2", "* * * * *", 2, 1, time.Now(), nil),
				}
				br.EXPECT().FindAll().Return(bs, nil)
				bob.EXPECT().SendBatchListResponse(bs)
			},
		}, {
			"normal scenario with query",
			args{"batch"},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				bs := []entity.Batch{
					newBatch(1, "batch", "server", "* * * * *", 2, 1, time.Now(), nil),
					newBatch(2, "batch2", "server2", "* * * * *", 2, 1, time.Now(), nil),
				}
				br.EXPECT().FindFilteredBy("batch").Return(bs, nil)
				bob.EXPECT().SendBatchListResponse(bs)
			},
		}, {
			"return 500 response",
			args{""},
			func(br *mock_core.MockBatchRepository, bob *mock_usecase.MockBatchOutputBoundary) {
				br.EXPECT().FindAll().Return([]entity.Batch{}, errors.New("sample error"))
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

			gblu := NewGetBatchListUsecase(br, bob)
			gblu.Execute(tt.args.query)
		})
	}
}

func newBatch(id int, batchName string, serverName string, cronSetting string, timeLimit int, estimationDuration int, startDate time.Time, endDate *time.Time) entity.Batch {
	b, _ := entity.NewBatch(id, batchName, serverName, cronSetting, timeLimit, estimationDuration, startDate, endDate)
	return *b
}
