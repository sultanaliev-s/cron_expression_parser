package cronparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	p, err := New("* * * * * cmd")

	assert.Nil(t, err, "should be nil")
	assert.NotEqual(t, p.minutes, make([]int, 0), "should not be empty")
	assert.NotEqual(t, p.hours, make([]int, 0), "should not be empty")
	assert.NotEqual(t, p.dayOfMonth, make([]int, 0), "should not be empty")
	assert.NotEqual(t, p.month, make([]int, 0), "should not be empty")
	assert.NotEqual(t, p.dayOfWeek, make([]int, 0), "should not be empty")
}

func TestParseMinutesEvery(t *testing.T) {
	p := Parser{}
	expected := make([]int, 0)
	for i := MinMinute; i <= MaxMinute; i++ {
		expected = append(expected, i)
	}

	err := p.parseMinutes("*")

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, p.minutes)
}

func TestParseHoursEvery(t *testing.T) {
	p := Parser{}
	expected := make([]int, 0)
	for i := MinHours; i <= MaxHours; i++ {
		expected = append(expected, i)
	}

	err := p.parseHours("*")

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, p.hours)
}

func TestParseDayOfMonthEvery(t *testing.T) {
	p := Parser{}
	expected := make([]int, 0)
	for i := MinDayOfMonth; i <= MaxDayOfMonth; i++ {
		expected = append(expected, i)
	}

	err := p.parseDayOfMonth("*")

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, p.dayOfMonth)
}

func TestParseMonthEvery(t *testing.T) {
	p := Parser{}
	expected := make([]int, 0)
	for i := MinMonth; i <= MaxMonth; i++ {
		expected = append(expected, i)
	}

	err := p.parseMonth("*")

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, p.month)
}

func TestParseDayOfWeekEvery(t *testing.T) {
	p := Parser{}
	expected := make([]int, 0)
	for i := MinDayOfWeek; i <= MaxDayOfWeek; i++ {
		expected = append(expected, i)
	}

	err := p.parseDayOfWeek("*")

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, p.dayOfWeek)
}

func TestParseEvery(t *testing.T) {
	p := Parser{}
	expected := make([]int, 60)
	for i := range expected {
		expected[i] = i
	}

	arr, err := p.parse("*", 0, 59)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, arr)
}

func TestParseRange(t *testing.T) {
	p := Parser{}
	expected := make([]int, 10)
	for i := range expected {
		expected[i] = i + 1
	}

	arr, err := p.parse("1-10", 0, 59)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, arr)
}

func TestParseList(t *testing.T) {
	p := Parser{}
	expected := []int{1, 10, 20, 30}

	arr, err := p.parse("1,10,20,30", 0, 59)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, arr)
}

func TestParseStep(t *testing.T) {
	p := Parser{}
	expected := []int{1, 11, 21, 31, 41, 51}

	arr, err := p.parse("1/10", 0, 59)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, arr)
}

func TestParseStepAsterisk(t *testing.T) {
	p := Parser{}
	expected := []int{0, 10, 20, 30, 40, 50}

	arr, err := p.parse("*/10", 0, 59)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, arr)
}

func TestParseNumber(t *testing.T) {
	p := Parser{}
	expected := []int{10}

	arr, err := p.parse("10", 0, 23)

	assert.Nil(t, err, "should be nil")
	assert.Equal(t, expected, arr)
}
