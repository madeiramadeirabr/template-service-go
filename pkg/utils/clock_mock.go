package utils

import "time"

type ClockMock struct{}

func (clock ClockMock) GetCurrentTimestamp() time.Time {
	return time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
}
