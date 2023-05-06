package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestHistoryRepositoryImpl_FindByIdAndBatchId(t *testing.T) {
	dialector := mysql.Open("root:root@tcp(127.0.0.1)/hottocoffee?parseTime=true&loc=Asia%2FTokyo")
	var db, _ = gorm.Open(dialector, &gorm.Config{})

	repo := NewHistoryRepositoryImpl(*db)
	repo.FindByIdAndBatchId(1, 1)
}
