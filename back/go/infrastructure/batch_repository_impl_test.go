package infrastructure_test

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HottoCoffee/HottoCoffee/core/entity"
	"github.com/HottoCoffee/HottoCoffee/infrastructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestBatchRepositoryImpl_FindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockDb, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	columns := []string{
		"id",
		"batch_name",
		"server_name",
		"cron_setting",
		"initial_date",
		"time_limit",
		"estimated_duration",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		recode  *infrastructure.BatchRecord
		want    *entity.Batch
		wantErr bool
	}{
		{
			"get 1 recode",
			fields{mockDb},
			args{1},
			&infrastructure.BatchRecord{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "name", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), 2, 1},
			&entity.Batch{Id: 1, BatchName: "name", ServerName: "server", CronSetting: newCronSetting("* * * * *"), TimeLimit: 2, EstimatedDuration: 1, StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), EndDate: nil},
			false,
		},
		{
			"get 0 recode",
			fields{mockDb},
			args{1},
			nil,
			nil,
			true,
		},
		{
			"get 1 broken recode",
			fields{mockDb},
			args{1},
			&infrastructure.BatchRecord{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), 2, 1},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := infrastructure.NewBatchRepository(tt.fields.db)

			if tt.recode == nil {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `batch` WHERE `batch`.`id` = ? AND `batch`.`deleted_at` IS NULL")).
					WithArgs(tt.args.id).
					WillReturnRows(sqlmock.NewRows(columns))
			} else {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `batch` WHERE `batch`.`id` = ? AND `batch`.`deleted_at` IS NULL")).
					WithArgs(tt.args.id).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow(tt.recode.ID, tt.recode.BatchName, tt.recode.ServerName, tt.recode.CronSetting, tt.recode.InitialDate, tt.recode.TimeLimit, tt.recode.EstimatedDuration, tt.recode.CreatedAt, tt.recode.UpdatedAt, tt.recode.DeletedAt))
			}

			got, err := br.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchRepositoryImpl.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchRepositoryImpl.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newCronSetting(v string) entity.CronSetting {
	s, _ := entity.NewCronSetting(v)
	return *s
}
