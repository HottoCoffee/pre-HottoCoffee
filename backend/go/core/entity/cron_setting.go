package entity

import (
	"time"

	"github.com/robfig/cron/v3"
)

type CronSetting struct {
	value    string
	schedule cron.Schedule
}

func NewCronSetting(v string) (*CronSetting, error) {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, err := p.Parse(v)
	if err != nil {
		return nil, NewDomainRuleViolationError("malformed schedule setting " + v)
	}

	return &CronSetting{value: v, schedule: s}, nil
}

func (cs CronSetting) ToString() string {
	return cs.value
}

func (cs CronSetting) ListSchedules(begin time.Time, endInclusive time.Time) (schedules []time.Time) {
	t := cs.schedule.Next(begin.Add(time.Minute * -1)) // Next doesn't return given time even though it matches cron settings.
	for t.Equal(endInclusive) || t.Before(endInclusive) {
		schedules = append(schedules, t)
		t = cs.schedule.Next(t)
	}
	return schedules
}
