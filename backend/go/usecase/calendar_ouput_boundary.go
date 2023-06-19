package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type CalendarOutputBoundary interface {
	SendCalendarResponse(calendarItems []entity.CalendarItem)
	SendInvalidRequestResponse(message string)
	SendInternalServerErrorResponse()
}
