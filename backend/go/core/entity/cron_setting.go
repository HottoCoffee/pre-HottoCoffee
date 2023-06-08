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

// Prev copy from https://github.com/robfig/cron/pull/361
func (cs CronSetting) Prev(t time.Time) time.Time {
	s := cs.schedule.(*cron.SpecSchedule)

	origLocation := t.Location()
	loc := s.Location
	if loc == time.Local {
		loc = t.Location()
	}
	if s.Location != time.Local {
		t = t.In(s.Location)
	}

	t = t.Add(-1*time.Second + time.Duration(t.Nanosecond())*time.Nanosecond)

	added := false

	yearLimit := t.Year() - 5

WRAP:
	if t.Year() < yearLimit {
		return time.Time{}
	}

	for 1<<uint(t.Month())&s.Month == 0 {
		cur := t.Month()
		if !added {
			added = true
		}
		t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc)
		t = t.Add(-1 * time.Second)

		if t.Month() > cur {
			goto WRAP
		}
	}

	for !dayMatches(s, t) {
		cur := t.Day()
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
			t = t.Add(-1 * time.Second)
		} else {
			t = t.AddDate(0, 0, -1)
		}
		if t.Day() > cur {
			goto WRAP
		}
	}

	for 1<<uint(t.Hour())&s.Hour == 0 {
		cur := t.Hour()
		if !added {
			added = true
			t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, loc)
			t = t.Add(-1 * time.Second)
		} else {
			t = t.Add(-1 * time.Hour)
		}
		if t.Hour() > cur {
			goto WRAP
		}
	}

	for 1<<uint(t.Minute())&s.Minute == 0 {
		cur := t.Minute()
		if !added {
			added = true
			t = t.Truncate(time.Minute)
			t = t.Add(-1 * time.Second)
		} else {
			t = t.Add(-1 * time.Minute)
		}

		if t.Minute() > cur {
			goto WRAP
		}
	}

	for 1<<uint(t.Second())&s.Second == 0 {
		cur := t.Second()
		if !added {
			added = true
			t = t.Truncate(time.Second)
		}
		t = t.Add(-1 * time.Second)

		if t.Second() > cur {
			goto WRAP
		}
	}

	return t.In(origLocation)
}

func dayMatches(s *cron.SpecSchedule, t time.Time) bool {
	var (
		domMatch = 1<<uint(t.Day())&s.Dom > 0
		dowMatch = 1<<uint(t.Weekday())&s.Dow > 0
	)
	starBit := uint64(1 << 63)
	if s.Dom&starBit > 0 || s.Dow&starBit > 0 {
		return domMatch && dowMatch
	}
	return domMatch || dowMatch
}
