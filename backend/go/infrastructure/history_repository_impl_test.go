package infrastructure

import (
	mock_core "github.com/HottoCoffee/HottoCoffee/.mock/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestHistoryRepositoryImpl_FindByIdAndBatchId(t *testing.T) {
	dialector := mysql.Open("root:root@tcp(127.0.0.1)/hottocoffee?parseTime=true&loc=Asia%2FTokyo")
	var db, _ = gorm.Open(dialector, &gorm.Config{})

	repo := NewHistoryRepositoryImpl(*db, mock_core.NewMockBatchRepository(nil))
	repo.FindAllDuring(time.Time{}, time.Time{})
}
