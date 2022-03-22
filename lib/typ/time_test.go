package typ

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Now().Format(TimeFormatDMGHIS))
	t.Log(time.Now().Format(TimeFormatDMG))
	t.Log(time.Now().Format(TimeFormatHIS))
}
