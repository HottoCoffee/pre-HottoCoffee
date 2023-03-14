package entity

import "time"

type Batch struct {
  Id int
  BatchName string
  ServerName string
  CronSetting string // TODO: use CronSetting-dedicated class
  TimeLimit int
  EsitimatedDuration int
  StartDate time.Time
  EndDate *time.Time
}
