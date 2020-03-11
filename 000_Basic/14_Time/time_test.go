package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToTimeStringRFC3339(t *testing.T) {
	assert := assert.New(t)
	result := time.Date(2020, time.October, 30, 23, 0, 0, 0, time.UTC)

	assert.Equal(ToTimeStringRFC3339(result), "2020-10-30T23:00:00+0000", "ToTimeStringRFC3339 failed")
}

func TestFromTimeStringRFC3339(t *testing.T) {
	assert := assert.New(t)

	result, err := FromTimeStringRFC3339("2020-10-30T23:00:00+0000")
	assert.Nil(err)
	assert.Equal(result.Year(), 2020, "TestToTimeString failed")
	assert.Equal(result.Month(), time.October, "TestToTimeString failed")
	assert.Equal(result.Day(), 30, "TestToTimeString failed")
	assert.Equal(result.Hour(), 23, "TestToTimeString failed")
}

func TestToTimeStringISO8601(t *testing.T) {
	assert := assert.New(t)
	result := time.Date(2020, time.October, 30, 23, 0, 0, 0, time.UTC)

	assert.Equal(ToTimeStringISO8601(result), "2020-10-30 23:00:00", "ToTimeStringISO8601 failed")
}

func TestFromTimeStringISO8601(t *testing.T) {
	assert := assert.New(t)

	result, err := FromTimeStringISO8601("2020-10-30 23:00:00")
	assert.Nil(err)
	assert.Equal(result.Year(), 2020, "TestToTimeString failed")
	assert.Equal(result.Month(), time.October, "TestToTimeString failed")
	assert.Equal(result.Day(), 30, "TestToTimeString failed")
	assert.Equal(result.Hour(), 23, "TestToTimeString failed")
}
