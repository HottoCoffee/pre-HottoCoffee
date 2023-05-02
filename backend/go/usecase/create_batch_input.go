package usecase

import "time"

type CreateBatchInput struct {
	BatchName   string    `json:"batch_name"`
	ServerName  string    `json:"server_name"`
	InitialDate time.Time `json:"initial_date"`
	TimeLimit   int       `json:"time_limit"`
	CronSetting string    `json:"cron_setting"`
}
