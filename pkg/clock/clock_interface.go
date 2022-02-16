package clock

import "time"

type Interface interface {
	GetCurrentTimestamp() time.Time
}
