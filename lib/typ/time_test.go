package typ

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Format(TimeFormatDMGHIS))
	fmt.Println(time.Now().Format(TimeFormatDMG))
	fmt.Println(time.Now().Format(TimeFormatHIS))
}
