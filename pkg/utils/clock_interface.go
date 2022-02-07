package utils

import "time"

type ClockInterface interface {
	GetCurrentTimestamp() time.Time
}
