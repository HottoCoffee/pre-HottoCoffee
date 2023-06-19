package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"time"
)

type GetCalendarUsecase struct {
	historyRepository      core.HistoryRepository
	calendarOutputBoundary CalendarOutputBoundary
}

func NewGetCalendarUsecase(repository core.HistoryRepository, boundary CalendarOutputBoundary) GetCalendarUsecase {
	return GetCalendarUsecase{
		historyRepository:      repository,
		calendarOutputBoundary: boundary,
	}
}

func (gcu GetCalendarUsecase) Execute(startDatetimeValue string, endDatetimeValue string) {
	startDatetime, err := time.Parse(time.RFC3339, startDatetimeValue)
	if err != nil {
		gcu.calendarOutputBoundary.SendInvalidRequestResponse(fmt.Sprintf("invalid start_datetime format. actual: %s", startDatetimeValue))
		return
	}
	endDatetime, err := time.Parse(time.RFC3339, endDatetimeValue)
	if err != nil {
		gcu.calendarOutputBoundary.SendInvalidRequestResponse(fmt.Sprintf("invalid end_datetime format. actual: %s", endDatetimeValue))
		return
	}

	batchExecutionHistoriesArray, err := gcu.historyRepository.FindAllDuring(startDatetime, endDatetime)
	if err != nil {
		fmt.Printf(err.Error())
		gcu.calendarOutputBoundary.SendInternalServerErrorResponse()
		return
	}

	calendarItems := core.GenerateCalendarItems(batchExecutionHistoriesArray, time.Now(), startDatetime, endDatetime)
	gcu.calendarOutputBoundary.SendCalendarResponse(calendarItems)
}
