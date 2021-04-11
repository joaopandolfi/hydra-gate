package format

import "time"

// CurrentDate time
func CurrentDate() time.Time {
	return time.Now().UTC()
}
