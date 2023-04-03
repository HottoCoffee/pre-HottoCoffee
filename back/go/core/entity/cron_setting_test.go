package entity

import (
	"testing"
	"time"
)

func TestNewCronSetting(t *testing.T) {
	_, err := NewCronSetting("0 * * * *")
	if err != nil {
		t.Error("err should not be null")
	}

	_, err2 := NewCronSetting("* * * *")
	if err2 == nil {
		t.Error("err should be null")
	}
}

func TestCronSetting_ListScedules(t *testing.T) {
	cs, _ := NewCronSetting("*/1 * * * *")
	ss := cs.ListScedules(time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local))
	if len(ss) != 61 {
		t.Errorf("Length of %ss should be 61", ss)
	}

	ss2 := cs.ListScedules(time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local))
	if !(len(ss2) == 1 && ss2[0] == time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)) {
		t.Errorf("%ss is not 2023/1/1 00:00:00", ss2[0])
	}
}
