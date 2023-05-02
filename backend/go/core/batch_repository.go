package core

import "github.com/HottoCoffee/HottoCoffee/core/entity"

type BatchRepository interface {
	FindById(id int) (*entity.Batch, error)
	FindAll() ([]entity.Batch, error)
	FindFilteredBy(query string) ([]entity.Batch, error)
	Save(batch entity.Batch) (*entity.Batch, error)
}
