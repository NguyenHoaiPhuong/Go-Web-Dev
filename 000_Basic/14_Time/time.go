package main

import (
	"time"
)

// ToTimeStringRFC3339 : convert time to string-format: "2006-01-02T15:04:05+0000"
func ToTimeStringRFC3339(t time.Time) string {
	// YYYY-MM-DD
	return t.Format("2006-01-02T15:04:05+0000")
}


// ToTimeStringISO8601 : convert time to string-format: "2006-01-02 15:04:05"
func ToTimeStringISO8601(t time.Time) string {
	// YYYY-MM-DD
	return t.Format("2006-01-02 15:04:05")
}

// FromTimeStringRFC3339 : convert string to time by format: "2006-01-02T15:04:05+0000"
func FromTimeStringRFC3339(timeStr string) (time.Time, error) {
	// YYYY-MM-DD
	layout := "2006-01-02T15:04:05+0000"
	result, err := time.Parse(layout, timeStr)
	return result, err
}

// FromTimeStringISO8601 : convert string to time by format: "2006-01-02 15:04:05"
func FromTimeStringISO8601(timeStr string) (time.Time, error) {
	// YYYY-MM-DD
	layout := "2006-01-02 15:04:05"
	result, err := time.Parse(layout, timeStr)
	return result, err
}