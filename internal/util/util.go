package util

import "time"

var Loc *time.Location

// Initialize the time location
func InitTimeZone() error {
	var err error
	Loc, err = time.LoadLocation("Asia/Jakarta")
	return err
}

func FormatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return d.String() // in microseconds
	} else if d < time.Second {
		return d.Truncate(time.Millisecond).String() // in milliseconds
	} else {
		return d.Truncate(time.Second).String() // in seconds
	}
}
