package test

import (
	"fmt"
	"testing"
	"time"

	"sungora/lib/typ"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Format(typ.TimeFormatDMGHIS))
	fmt.Println(time.Now().Format(typ.TimeFormatDMG))
	fmt.Println(time.Now().Format(typ.TimeFormatHIS))
}
