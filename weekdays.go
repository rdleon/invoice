package main

import (
	"math"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func weekdayOffset(d time.Time) int {
	wd := d.Weekday()
	if wd == time.Sunday {
		return 6
	}
	return int(wd) - 1
}

func weekdays(t time.Time) int {
	// WARN: Doesn't take holidays into account

	// Copied from https://stackoverflow.com/questions/31327124/how-to-calculate-number-of-business-days-in-golang
	var firstDay = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	var lastDay = time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	startOffset := weekdayOffset(firstDay)
	startTime := firstDay.AddDate(0, 0, -startOffset)

	endOffset := weekdayOffset(lastDay)
	endTime := lastDay.AddDate(0, 0, -endOffset)

	diff := endTime.Sub(startTime)

	weeks := int(math.Round((diff.Hours() / 24) / 7))

	days := min(endOffset, 5) - min(startOffset, 5)

	return (weeks * 5) + days
}
