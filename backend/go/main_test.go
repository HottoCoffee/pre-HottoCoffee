package main

import (
	"encoding/json"
	"github.com/HottoCoffee/HottoCoffee/infrastructure"
	"github.com/HottoCoffee/HottoCoffee/util"
	"github.com/google/go-cmp/cmp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

var dialector = mysql.Open("root:root@tcp(127.0.0.1)/hottocoffee?parseTime=true&loc=Asia%2FTokyo")
var db, _ = gorm.Open(dialector, &gorm.Config{})

func truncate() {
	db.Exec("set foreign_key_checks = 0")
	db.Exec("truncate table batch")
	db.Exec("truncate table history")
	db.Exec("set foreign_key_checks = 1")
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	resRecorder := httptest.NewRecorder()
	r.ServeHTTP(resRecorder, req)
	return resRecorder
}

func TestGetBatchIdApi(t *testing.T) {
	// setup
	if testing.Short() {
		t.Skip()
	}
	truncate()

	route := SetUp()

	// given
	batchRecord := infrastructure.BatchRecord{BatchName: "hoge", ServerName: "fuga", CronSetting: "0 * * * *", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), TimeLimit: 100, EstimatedDuration: 50}
	batchRecord.ID = 1
	db.Create(&batchRecord)

	// when
	resRecorder := performRequest(route, "GET", "/api/batch/1", nil)

	var got interface{}
	_ = json.Unmarshal(resRecorder.Body.Bytes(), &got)

	// then
	want := map[string]interface{}{
		"id":           float64(1),
		"batch_name":   "hoge",
		"server_name":  "fuga",
		"cron_setting": "0 * * * *",
		"initial_date": "2023-01-01T00:00:00+09:00",
		"time_limit":   float64(100),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GET /batch/{id} got = %v, want %v", got, want)
	}
}

func TestGetBatchListApi(t *testing.T) {
	// setup
	if testing.Short() {
		t.Skip()
	}
	truncate()

	route := SetUp()

	// given
	batchRecords := []infrastructure.BatchRecord{
		{BatchName: "hoge", ServerName: "fuga", CronSetting: "0 * * * *", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), TimeLimit: 100, EstimatedDuration: 50},
		{BatchName: "piyo", ServerName: "foo", CronSetting: "0 * * * *", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, util.JST), TimeLimit: 100, EstimatedDuration: 50},
	}
	batchRecords[0].ID = 1
	batchRecords[1].ID = 2
	db.Create(batchRecords)

	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"get all batches without query",
			args{"/api/batch"},
			[]any{
				map[string]interface{}{
					"id":           float64(1),
					"batch_name":   "hoge",
					"server_name":  "fuga",
					"cron_setting": "0 * * * *",
					"initial_date": "2023-01-01T00:00:00+09:00",
					"time_limit":   float64(100),
				},
				map[string]interface{}{
					"id":           float64(2),
					"batch_name":   "piyo",
					"server_name":  "foo",
					"cron_setting": "0 * * * *",
					"initial_date": "2023-01-01T00:00:00+09:00",
					"time_limit":   float64(100),
				},
			},
		}, {
			"get batches by query",
			args{"/api/batch?query=hog"},
			[]any{
				map[string]any{
					"id":           float64(1),
					"batch_name":   "hoge",
					"server_name":  "fuga",
					"cron_setting": "0 * * * *",
					"initial_date": "2023-01-01T00:00:00+09:00",
					"time_limit":   float64(100),
				},
			},
		}, {
			"get empty batch list",
			args{"/api/batch?query=not_exist"},
			[]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resRecorder := performRequest(route, "GET", tt.args.url, nil)
			var got interface{}
			_ = json.Unmarshal(resRecorder.Body.Bytes(), &got)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("GET %v got = %v, want %v, diff %v", tt.args.url, got, tt.want, diff)
			}
		})
	}
}

func TestPostBatchApi(t *testing.T) {
	// setup
	if testing.Short() {
		t.Skip()
	}
	truncate()

	route := SetUp()

	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"post valid batch",
			args{`{"batch_name":"hoge","server_name":"fuga","cron_setting":"0 * * * *", "initial_date":"2023-01-01T00:00:00+09:00","time_limit":1}`},
			map[string]interface{}{
				"id":           float64(1),
				"batch_name":   "hoge",
				"server_name":  "fuga",
				"cron_setting": "0 * * * *",
				"initial_date": "2023-01-01T00:00:00+09:00",
				"time_limit":   float64(1),
			},
		}, {
			"post invalid batch",
			args{`{}`},
			map[string]interface{}{
				"status":  float64(400),
				"message": "batch name should not be empty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resRecorder := performRequest(route, "POST", "/api/batch", strings.NewReader(tt.args.body))
			var got interface{}
			_ = json.Unmarshal(resRecorder.Body.Bytes(), &got)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("POST /api/batch got = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}
