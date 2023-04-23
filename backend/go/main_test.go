package main

import (
	"encoding/json"
	"github.com/HottoCoffee/HottoCoffee/infrastructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
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
	batchRecord := infrastructure.BatchRecord{BatchName: "hoge", ServerName: "fuga", CronSetting: "0 * * * *", InitialDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), TimeLimit: 100, EstimatedDuration: 50}
	batchRecord.ID = 1
	db.Create(&batchRecord)

	// when
	resRecorder := performRequest(route, "GET", "/api/batch/1", nil)

	var got map[string]any
	json.Unmarshal(resRecorder.Body.Bytes(), &got)

	// then
	want := map[string]interface{}{
		"id":           1,
		"batch_name":   "hoge",
		"server_name":  "fuga",
		"cron_setting": "0 * * * *",
		"initial_date": "2023-01-01T00:00:00Z",
		"time_limit":   100,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GET /batch/{id} got = %v, want %v", got, want)
	}
}
