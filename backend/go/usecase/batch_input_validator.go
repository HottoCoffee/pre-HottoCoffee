package usecase

import "github.com/HottoCoffee/HottoCoffee/core/entity"

func validateBatchInput(input BatchInput) error {
	dummyId := 1
	if _, err := entity.NewBatch(dummyId, input.BatchName, input.ServerName, input.CronSetting, input.TimeLimit, 0, input.InitialDate, nil); err != nil {
		return err
	}
	return nil
}
