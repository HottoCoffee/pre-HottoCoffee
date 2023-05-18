package usecase

import (
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"time"
)

type GetCalendarUsecase struct {
	historyOutputBoundary HistoryOutputBoundary
	historyRepository     core.HistoryRepository
}

func NewGetCalendarUsecase(historyOutputBoundary HistoryOutputBoundary, historyRepository core.HistoryRepository) GetCalendarUsecase {
	return GetCalendarUsecase{
		historyOutputBoundary: historyOutputBoundary,
		historyRepository:     historyRepository,
	}
}

func (gcu GetCalendarUsecase) Execute(startDatetimeInput string, endDatetimeInput string) {
	startDatetime, err1 := time.Parse(time.RFC3339Nano, startDatetimeInput)
	endDatetime, err2 := time.Parse(time.RFC3339Nano, endDatetimeInput)
	if err1 != nil || err2 != nil {
		gcu.historyOutputBoundary.SendBadRequestResponse("input should be datetime format")
		return
	}

	historiesArray, err := gcu.historyRepository.FindAllDuring(startDatetime, endDatetime)
	if err != nil {
		gcu.historyOutputBoundary.SendInternalServerErrorResponse()
	}
	var histories []entity.History
	for _, h := range historiesArray {
		for _, s := range h.Batch.CronSetting.ListSchedules(startDatetime, endDatetime) {
			histories = append(histories, h.GetHistoryExecutedOn(s))
		}
	}
	gcu.historyOutputBoundary.SendCalenderResponse(histories)
}
