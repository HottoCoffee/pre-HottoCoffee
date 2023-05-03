package infrastructure_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/util"
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
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
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
		record  *infrastructure.BatchRecord
		want    *entity.Batch
		wantErr bool
	}{
		{
			"get 1 record",
			fields{mockDb},
			args{1},
			&infrastructure.BatchRecord{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "name", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), 2, 1},
			&entity.Batch{Id: 1, BatchName: "name", ServerName: "server", CronSetting: newCronSetting("* * * * *"), TimeLimit: 2, EstimatedDuration: 1, StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), EndDate: nil},
			false,
		},
		{
			"get 0 record",
			fields{mockDb},
			args{1},
			nil,
			nil,
			true,
		},
		{
			"get 1 broken record",
			fields{mockDb},
			args{1},
			&infrastructure.BatchRecord{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), 2, 1},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := infrastructure.NewBatchRepository(tt.fields.db)

			if tt.record == nil {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `batch` WHERE `batch`.`id` = ? AND `batch`.`deleted_at` IS NULL")).
					WithArgs(tt.args.id).
					WillReturnRows(sqlmock.NewRows(columns))
			} else {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `batch` WHERE `batch`.`id` = ? AND `batch`.`deleted_at` IS NULL")).
					WithArgs(tt.args.id).
					WillReturnRows(sqlmock.NewRows(columns).
						AddRow(tt.record.ID, tt.record.BatchName, tt.record.ServerName, tt.record.CronSetting, tt.record.InitialDate, tt.record.TimeLimit, tt.record.EstimatedDuration, tt.record.CreatedAt, tt.record.UpdatedAt, tt.record.DeletedAt))
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

func TestBatchRepositoryImpl_FindAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
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
	tests := []struct {
		name    string
		fields  fields
		records []infrastructure.BatchRecord
		want    []entity.Batch
		wantErr bool
	}{
		{
			"get records",
			fields{mockDb},
			[]infrastructure.BatchRecord{
				{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "name", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), 2, 1},
				{gorm.Model{ID: 2, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "name2", "server2", "* * * * *", time.Date(2023, 2, 3, 4, 5, 6, 7, util.JST), 4, 3},
			},
			[]entity.Batch{
				{Id: 1, BatchName: "name", ServerName: "server", CronSetting: newCronSetting("* * * * *"), TimeLimit: 2, EstimatedDuration: 1, StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), EndDate: nil},
				{Id: 2, BatchName: "name2", ServerName: "server2", CronSetting: newCronSetting("* * * * *"), TimeLimit: 4, EstimatedDuration: 3, StartDate: time.Date(2023, 2, 3, 4, 5, 6, 7, util.JST), EndDate: nil},
			},
			false,
		},
		{
			"get 0 records",
			fields{mockDb},
			[]infrastructure.BatchRecord{},
			[]entity.Batch{},
			false,
		},
		{
			"get broken records",
			fields{mockDb},
			[]infrastructure.BatchRecord{{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), 2, 1}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := infrastructure.NewBatchRepository(tt.fields.db)

			rows := sqlmock.NewRows(columns)
			for i := range tt.records {
				record := tt.records[i]
				rows.AddRow(record.ID, record.BatchName, record.ServerName, record.CronSetting, record.InitialDate, record.TimeLimit, record.EstimatedDuration, record.CreatedAt, record.UpdatedAt, record.DeletedAt)
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `batch` WHERE `batch`.`deleted_at` IS NULL")).
				WillReturnRows(rows)

			got, err := br.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchRepositoryImpl.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.want) == 0 {
				if len(got) != 0 {
					t.Errorf("BatchRepositoryImpl.FindById() = %v, want %v", got, tt.want)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchRepositoryImpl.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchRepositoryImpl_FindFilteredBy(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
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
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		records []infrastructure.BatchRecord
		want    []entity.Batch
		wantErr bool
	}{
		{
			"get all records",
			fields{mockDb},
			args{"name"},
			[]infrastructure.BatchRecord{
				{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "name", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), 2, 1},
				{gorm.Model{ID: 2, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "name2", "server2", "* * * * *", time.Date(2023, 2, 3, 4, 5, 6, 7, util.JST), 4, 3},
			},
			[]entity.Batch{
				{Id: 1, BatchName: "name", ServerName: "server", CronSetting: newCronSetting("* * * * *"), TimeLimit: 2, EstimatedDuration: 1, StartDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), EndDate: nil},
				{Id: 2, BatchName: "name2", ServerName: "server2", CronSetting: newCronSetting("* * * * *"), TimeLimit: 4, EstimatedDuration: 3, StartDate: time.Date(2023, 2, 3, 4, 5, 6, 7, util.JST), EndDate: nil},
			},
			false,
		},
		{
			"get 0 records",
			fields{mockDb},
			args{"non_exist_name"},
			[]infrastructure.BatchRecord{},
			[]entity.Batch{},
			false,
		},
		{
			"get broken records",
			fields{mockDb},
			args{"name"},
			[]infrastructure.BatchRecord{{gorm.Model{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: false}}, "", "server", "* * * * *", time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), 2, 1}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := infrastructure.NewBatchRepository(tt.fields.db)

			rows := sqlmock.NewRows(columns)
			for i := range tt.records {
				record := tt.records[i]
				rows.AddRow(record.ID, record.BatchName, record.ServerName, record.CronSetting, record.InitialDate, record.TimeLimit, record.EstimatedDuration, record.CreatedAt, record.UpdatedAt, record.DeletedAt)
			}

			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `batch` WHERE (batch_name LIKE ? OR server_name LIKE ? OR cron_setting LIKE ?) AND `batch`.`deleted_at` IS NULL")).
				WithArgs("%"+tt.args.query+"%", "%"+tt.args.query+"%", "%"+tt.args.query+"%").
				WillReturnRows(rows)

			got, err := br.FindFilteredBy(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchRepositoryImpl.FindFilteredBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.want) == 0 {
				if len(got) != 0 {
					t.Errorf("BatchRepositoryImpl.FindFilteredBy() = %v, want %v", got, tt.want)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BatchRepositoryImpl.FindFilteredBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchRepositoryImpl_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	mockDb, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})

	type args struct {
		batchName   string
		serverName  string
		cronSetting string
		timeLimit   int
		startDate   time.Time
	}
	tests := []struct {
		name       string
		args       args
		sqlResult  sql.Result
		want       *entity.Batch
		wantErr    bool
		wantErrMsg string
	}{
		{
			"successful insertion",
			args{"batch", "server", "* * * * *", 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST)},
			sqlmock.NewResult(1, 1),
			newBatch(1, "batch", "server", "* * * * *", 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST)),
			false,
			"",
		}, {
			"failure insertion",
			args{"batch", "server", "* * * * *", 1, time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST)},
			sqlmock.NewResult(0, 0),
			nil,
			true,
			`failed to insert batch record: \{\{0 .+ \{0001-01-01 00:00:00 \+0000 UTC false\}\} batch server \* \* \* \* \* 2023-01-01 00:00:00 \+0900 JST 1 0\}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := infrastructure.NewBatchRepository(mockDb)

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `batch` (`created_at`,`updated_at`,`deleted_at`,`batch_name`,`server_name`,`cron_setting`,`initial_date`,`time_limit`,`estimated_duration`) VALUES (?,?,?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, tt.args.batchName, tt.args.serverName, tt.args.cronSetting, tt.args.startDate, tt.args.timeLimit, 0).
				WillReturnResult(tt.sqlResult)
			mock.ExpectCommit()

			got, err := br.Create(tt.args.batchName, tt.args.serverName, tt.args.cronSetting, tt.args.timeLimit, tt.args.startDate)
			if tt.wantErr {
				compile := regexp.MustCompile(tt.wantErrMsg)
				fmt.Println(compile)
				if err == nil {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				} else if !compile.MatchString(err.Error()) {
					t.Errorf("Create() error = %v, wantErrMsg %v", err.Error(), tt.wantErrMsg)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func newBatch(id int, batchName string, serverName string, cronSetting string, timeLimit int, startDate time.Time) *entity.Batch {
	batch, _ := entity.NewBatch(id, batchName, serverName, cronSetting, timeLimit, 0, startDate, nil)
	return batch
}

func newCronSetting(v string) entity.CronSetting {
	s, _ := entity.NewCronSetting(v)
	return *s
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
