package CoreModels

import (
	"encoding/json"
	"time"
)

// Constant
const (
	DateLayoutISO        = "2006-01-02T15:04:05.000Z"
	DateStandardResponse = "2006-01-02T15:04:05.000Z07:00"
)

// TimeResponse ...
type TimeResponse struct {
	Time time.Time
}

// MarshalJSON ...
func (t TimeResponse) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(t.Time.Format(DateStandardResponse))
}

// IsBefore ...
func (t *TimeResponse) IsBefore(c TimeResponse) bool {
	return t.Time.Before(c.Time)
}

// IsAfter ...
func (t *TimeResponse) IsAfter(c TimeResponse) bool {
	return t.Time.After(c.Time)
}

// TimeResponseInit ...
func TimeResponseInit(time time.Time) TimeResponse {
	return TimeResponse{
		time,
	}
}

// TimeResponseNow ...
func TimeResponseNow() TimeResponse {
	return TimeResponseInit(time.Now())
}

// TimeResponseCustomFromISOString ...
func TimeResponseCustomFromISOString(timeString string) TimeResponse {
	t, _ := time.Parse(DateLayoutISO, timeString)
	return TimeResponseInit(t)
}
