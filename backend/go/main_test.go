package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1)/hottocoffee?parseTime=true&loc=Asia%2FTokyo")

func truncate() {
	_, _ = db.Exec("set foreign_key_checks = 0")
	_, _ = db.Exec("truncate table batch")
	_, _ = db.Exec("truncate table history")
	_, _ = db.Exec("set foreign_key_checks = 1")
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
	_, _ = db.Exec("insert into batch (id, batch_name, server_name, cron_setting, initial_date, time_limit, estimated_duration) values (1, 'hoge', 'fuga', '0 * * * *', '2023-01-01', 100, 50)")

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
	_, _ = db.Exec("insert into batch (id, batch_name, server_name, cron_setting, initial_date, time_limit, estimated_duration) values (1, 'hoge', 'fuga', '0 * * * *', '2023-01-01', 100, 50), (2, 'piyo', 'foo', '0 * * * *', '2023-01-01', 100, 50)")

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
		truncate()
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

func TestPutBatchApi(t *testing.T) {
	// setup
	if testing.Short() {
		t.Skip()
	}

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
			"put valid batch",
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
			"put invalid batch",
			args{`{}`},
			map[string]interface{}{
				"status":  float64(400),
				"message": "batch name should not be empty",
			},
		},
	}
	for _, tt := range tests {
		truncate()
		_, _ = db.Exec("insert into batch (batch_name, server_name, cron_setting, initial_date, time_limit, estimated_duration) values ('batch', 'server', '1 * * * *', '2023-04-01', 100, 0)")
		t.Run(tt.name, func(t *testing.T) {
			resRecorder := performRequest(route, "PUT", "/api/batch/1", strings.NewReader(tt.args.body))
			var got interface{}
			_ = json.Unmarshal(resRecorder.Body.Bytes(), &got)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("POST /api/batch got = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}

func TestGetHistoryApi(t *testing.T) {
	// setup
	if testing.Short() {
		t.Skip()
	}
	truncate()

	route := SetUp()

	// given
	_, _ = db.Exec(`insert into batch(id, batch_name, server_name, cron_setting, initial_date, time_limit, estimated_duration) values (1, 'hoge', 'fuga', '0 * * * *', '2023-01-01', 100, 50)`)
	_, _ = db.Exec(`insert into history(id, batch_id, status, start_datetime, finish_datetime) values (1, 1, 'success', '2023-01-01 01:00:00', '2023-01-01 01:01:00'), (2, 1, 'failed', '2023-01-01 02:00:00', '2023-01-01 02:01:00')`)

	type args struct {
		batchId   string
		historyId string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"get success history",
			args{batchId: "1", historyId: "1"},
			map[string]interface{}{
				"history_id":      float64(1),
				"batch_id":        float64(1),
				"batch_name":      "hoge",
				"start_datetime":  "2023-01-01T01:00:00+09:00",
				"finish_datetime": "2023-01-01T01:01:00+09:00",
				"status":          "success",
			},
		},
		{
			"get failed history",
			args{batchId: "1", historyId: "2"},
			map[string]interface{}{
				"history_id":      float64(2),
				"batch_id":        float64(1),
				"batch_name":      "hoge",
				"start_datetime":  "2023-01-01T02:00:00+09:00",
				"finish_datetime": "2023-01-01T02:01:00+09:00",
				"status":          "failed",
			},
		},
		{
			"not found response",
			args{batchId: "1", historyId: "3"},
			map[string]interface{}{
				"status":  float64(404),
				"message": "Not Found",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resRecorder := performRequest(route, "GET", fmt.Sprintf("/api/batch/%v/history/%v", tt.args.batchId, tt.args.historyId), nil)
			var got interface{}
			_ = json.Unmarshal(resRecorder.Body.Bytes(), &got)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("POST /api/batch got = %v, want %v, diff %v", got, tt.want, diff)
			}
		})
	}
}
