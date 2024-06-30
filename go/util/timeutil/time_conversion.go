package timeutil

import "time"

func TimeToDateString(targetDate *time.Time) string {
	if targetDate != nil {
		return targetDate.Format(time.DateOnly)
	}
	return ""
}

func StringToDate(timeString string) *time.Time {
	if targetDate, err := time.Parse(time.DateOnly, timeString); err == nil {
		return &targetDate
	}
	return nil
}
