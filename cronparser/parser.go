package cronparser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MinMinute     = 0
	MaxMinute     = 59
	MinHours      = 0
	MaxHours      = 23
	MinDayOfMonth = 1
	MaxDayOfMonth = 31
	MinMonth      = 1
	MaxMonth      = 12
	MinDayOfWeek  = 1
	MaxDayOfWeek  = 7
)

type Parser struct {
	minutes    []int
	hours      []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
	command    string
}

func New(expression string) (Parser, error) {
	p := Parser{}

	expression = strings.Trim(expression, " ")
	parts := strings.Split(expression, " ")
	if len(parts) != 6 {
		return p, fmt.Errorf("invalid amount of fields, got %d", len(parts))
	}

	err := p.parseMinutes(parts[0])
	if err != nil {
		return p, err
	}
	err = p.parseHours(parts[1])
	if err != nil {
		return p, err
	}
	err = p.parseDayOfMonth(parts[2])
	if err != nil {
		return p, err
	}
	err = p.parseMonth(parts[3])
	if err != nil {
		return p, err
	}
	err = p.parseDayOfWeek(parts[4])
	if err != nil {
		return p, err
	}
	p.command = parts[5]

	return p, nil
}

func (p *Parser) parseMinutes(exp string) error {
	arr, err := p.parse(exp, MinMinute, MaxMinute)
	if err != nil {
		return err
	}
	p.minutes = arr
	return nil
}

func (p *Parser) parseHours(exp string) error {
	arr, err := p.parse(exp, MinHours, MaxHours)
	if err != nil {
		return err
	}
	p.hours = arr
	return nil
}

func (p *Parser) parseDayOfMonth(exp string) error {
	arr, err := p.parse(exp, MinDayOfMonth, MaxDayOfMonth)
	if err != nil {
		return err
	}
	p.dayOfMonth = arr
	return nil
}

func (p *Parser) parseMonth(exp string) error {
	arr, err := p.parse(exp, MinMonth, MaxMonth)
	if err != nil {
		return err
	}
	p.month = arr
	return nil
}

func (p *Parser) parseDayOfWeek(exp string) error {
	arr, err := p.parse(exp, MinDayOfWeek, MaxDayOfWeek)
	if err != nil {
		return err
	}
	p.dayOfWeek = arr
	return nil
}

func (p *Parser) parse(exp string, min, max int) ([]int, error) {
	arr := make([]int, 0)
	if exp == "*" {
		arr = make([]int, max-min+1)
		for i := range arr {
			arr[i] = min + i
		}
		return arr, nil
	}

	parts := strings.Split(exp, "-")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid range expression")
	}
	if len(parts) == 2 {
		if parts[0] == "" || parts[1] == "" {
			return nil, fmt.Errorf("invalid range expression")
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		if !isInRange(start, min, max) {
			return nil, fmt.Errorf("outside the range of a minute")
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		if !isInRange(end, min, max) {
			return nil, fmt.Errorf("outside the range of a minute")
		}

		arr = make([]int, end-start+1)
		for i, j := start, 0; i <= end; i, j = i+1, j+1 {
			arr[j] = i
		}
		return arr, nil
	}

	parts = strings.Split(exp, ",")
	if len(parts) > 1 {
		arr = make([]int, len(parts))
		for i := range parts {
			if parts[i] == "" {
				return nil, fmt.Errorf("invalid list expression")
			}
			min, err := strconv.Atoi(parts[i])
			if err != nil {
				return nil, nil
			}
			arr[i] = min
		}
		return arr, nil
	}

	parts = strings.Split(exp, "/")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid step expression")
	}
	if len(parts) == 2 {
		start := 0
		if parts[0] != "*" {
			min, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}
			start = min
		}

		step, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		arr = append(arr, start)
		start += step
		for start < max {
			arr = append(arr, start)
			start += step
		}
		return arr, nil
	}

	min, err := strconv.Atoi(exp)
	if err != nil {
		return nil, err
	}
	arr = append(arr, min)

	return arr, nil
}

func isInRange(start, min, max int) bool {
	return min <= start && start <= max
}

func (p Parser) String() string {
	var minutes []string
	var hours []string
	var dayOfMonth []string
	var month []string
	var dayOfWeek []string
	for i := range p.minutes {
		minutes = append(minutes, strconv.Itoa(p.minutes[i]))
	}
	for i := range p.hours {
		hours = append(hours, strconv.Itoa(p.hours[i]))
	}
	for i := range p.dayOfMonth {
		dayOfMonth = append(dayOfMonth, strconv.Itoa(p.dayOfMonth[i]))
	}
	for i := range p.month {
		month = append(month, strconv.Itoa(p.month[i]))
	}
	for i := range p.dayOfWeek {
		dayOfWeek = append(dayOfWeek, strconv.Itoa(p.dayOfWeek[i]))
	}
	res := fmt.Sprintf(
		"minute %s\nhour %s\nday of month %s\nmonth %s\nday of week %s\ncommand %s\n",
		strings.Join(minutes, " "),
		strings.Join(hours, " "),
		strings.Join(dayOfMonth, " "),
		strings.Join(month, " "),
		strings.Join(dayOfWeek, " "),
		p.command,
	)
	return res
}
