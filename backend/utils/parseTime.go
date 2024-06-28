package utils

import (
	"time"
)

func ParseTime(timeString string) (time.Time, error) {
	layout := "15:04"
	parsedTime, err := time.Parse(layout, timeString)

	// Parsing time with only hour and minute specified will make the year to be 0, add 1 year to avoid error in gorm
	parsedTime = parsedTime.AddDate(1, 0, 0)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
