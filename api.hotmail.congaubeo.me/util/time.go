package util

import "time"

// Constant
const (
	DateISOFormat = "2006-01-02T15:04:05.000Z"
)

// TimeParseISODate ...
func TimeParseISODate(value string) time.Time {
	t, _ := TimeHCMFormatToLayout(value, DateISOFormat)
	return t
}

// TimeHCMFormatToLayout ...
func TimeHCMFormatToLayout(str string, layout string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	t, err := time.ParseInLocation(layout, str, loc)
	return t, err
}

// TimeCheckTwoDatesEqual ...
func TimeCheckTwoDatesEqual(date1 time.Time, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// TimeParseISODateAndChangeToStartOfDate ...
func TimeParseISODateAndChangeToStartOfDate(value string) time.Time {
	t, _ := TimeHCMFormatToLayout(value, DateISOFormat)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// TimeParseISODateAndChangeToEndOfDate ...
func TimeParseISODateAndChangeToEndOfDate(value string) time.Time {
	t, _ := TimeHCMFormatToLayout(value, DateISOFormat)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// TimeStartOfToday ...
func TimeStartOfToday() time.Time {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

// TimeEndOfToday ...
func TimeEndOfToday() time.Time {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
}
