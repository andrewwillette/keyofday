package key

import (
	"fmt"
	"testing"
	"time"
)

func TestGetKeyOfDay(t *testing.T) {
	keyofday := TodaysKey()
	tomorrowsKey := Key(time.Now().AddDate(0, 0, 1))
	yesterdaysKey := Key(time.Now().AddDate(0, 0, -1))
	fmt.Println(yesterdaysKey)
	fmt.Println(keyofday)
	fmt.Println(tomorrowsKey)
}
