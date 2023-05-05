package usecase

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/core"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"strconv"
)

type ChangeBatchUsecase struct {
	br  core.BatchRepository
	bob BatchOutputBoundary
}

func NewChangeBatchUsecase(br core.BatchRepository, bob BatchOutputBoundary) ChangeBatchUsecase {
	return ChangeBatchUsecase{br: br, bob: bob}
}

func (cbu ChangeBatchUsecase) Execute(stringId string, input BatchInput) {
	id, err := strconv.Atoi(stringId)
	if err != nil {
		cbu.bob.SendNotFoundResponse()
		return
	}

	b, err := cbu.br.FindById(id)
	if err != nil {
		cbu.bob.SendNotFoundResponse()
		return
	}

	err = validateBatchInput(input)
	if err != nil {
		cbu.bob.SendInvalidRequestResponse(err.Error())
		return
	}

	b.BatchName = input.BatchName
	b.ServerName = input.ServerName
	cronSetting, _ := entity.NewCronSetting(input.CronSetting)
	b.CronSetting = *cronSetting
	b.TimeLimit = input.TimeLimit
	b.StartDate = input.InitialDate

	err = cbu.br.Save(*b)
	if err != nil {
		_ = fmt.Errorf(err.Error())
		cbu.bob.SendInternalServerErrorResponse()
		return
	}

	cbu.bob.SendBatchResponse(*b)
}
