package core

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type BatchRepository interface {
	FindById(id int) (*entity.Batch, error)
}
