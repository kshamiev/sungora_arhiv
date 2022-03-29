package conf

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Now().Format(TimeFormatGMDHIS))
	t.Log(time.Now().Format(TimeFormatDMGHIS))
	t.Log(time.Now().Format(TimeFormatDMG))
	t.Log(time.Now().Format(TimeFormatHIS))

	loc, _ := time.LoadLocation("Europe/Moscow")
	t.Log(time.Now().In(loc).Format(TimeFormatGMDHIS))
}
