package main

import (
	"testing"
	"time"
)

func TestWorkDays(t *testing.T) {
	type test struct {
		input time.Time
		days  int
	}

	tests := []test{
		{time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC), 20},
		{time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC), 21},
		{time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC), 23},
		{time.Date(2024, 7, 8, 0, 0, 0, 0, time.UTC), 23},
		{time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC), 20},
		{time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC), 23},
	}

	for _, tc := range tests {
		got := workDays(tc.input)
		if got != tc.days {
			t.Fatalf("Got %d days for %v, wanted %d", got, tc.input, tc.days)
		}
	}

}
