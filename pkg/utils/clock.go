package utils

import "time"

type Clock struct{}

func (clock Clock) GetCurrentTimestamp() time.Time {
	return time.Now()
}
