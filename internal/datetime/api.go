package datetime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Parser struct {
	now *time.Time
}

type ParserType = func(exp string) (time.Time, error)

type Option func(p *Parser)

func WithNow(now time.Time) Option {
	return func(p *Parser) {
		p.now = &now
	}
}

func NewParser(opts ...Option) *Parser {
	p := &Parser{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func DateOnly(t time.Time) time.Time {
	y, m, d := t.Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return date
}

func (p Parser) ParseRelativeDay(exp string) (time.Time, error) {

	// today
	if exp == "-" {
		return p.getNow(), nil
	}

	// negative offset
	if strings.HasSuffix(exp, "-") {
		offsetRaw := strings.TrimSuffix(exp, "-")
		offset, err := strconv.Atoi(offsetRaw)
		if err != nil {
			return time.Time{}, err
		}
		return p.getNow().Add(time.Duration(-offset) * time.Hour * 24), err
	}

	// positive offset
	if strings.HasSuffix(exp, "+") {
		offsetRaw := strings.TrimSuffix(exp, "+")
		offset, err := strconv.Atoi(offsetRaw)
		if err != nil {
			return time.Time{}, err
		}
		return p.getNow().Add(time.Duration(offset) * time.Hour * 24), err
	}

	return time.Time{}, fmt.Errorf("invalid date expression: %s", exp)

}

func (p Parser) ParseRelativeWeek(exp string) (time.Time, error) {

	// today
	if exp == "-" {
		return p.getNow(), nil
	}

	// negative offset
	if strings.HasSuffix(exp, "-") {
		offsetRaw := strings.TrimSuffix(exp, "-")
		offset, err := strconv.Atoi(offsetRaw)
		if err != nil {
			return time.Time{}, err
		}
		return p.getNow().Add(time.Duration(-offset) * time.Hour * 24 * 7), err
	}

	// positive offset
	if strings.HasSuffix(exp, "+") {
		offsetRaw := strings.TrimSuffix(exp, "+")
		offset, err := strconv.Atoi(offsetRaw)
		if err != nil {
			return time.Time{}, err
		}
		return p.getNow().Add(time.Duration(offset) * time.Hour * 24 * 7), err
	}

	return time.Time{}, fmt.Errorf("invalid date expression: %s", exp)
}

func (p Parser) ParseRelativeMonth(exp string) (time.Time, error) {

	// current month
	if exp == "-" {
		return p.getNow(), nil
	}

	// negative offset
	if strings.HasSuffix(exp, "-") {
		offsetRaw := strings.TrimSuffix(exp, "-")
		offset, err := strconv.Atoi(offsetRaw)
		if err != nil {
			return time.Time{}, err
		}
		y, m, _ := p.getNow().Date()
		m -= time.Month(offset)
		for {
			if m > 0 {
				break
			}
			y -= 1
			m += 12
		}
		newMonth := time.Date(y, m, 1, 0, 0, 0, 0, p.getNow().Location())
		return newMonth, nil
	}

	return time.Time{}, fmt.Errorf("invalid date expression: %s", exp)
}

func (p Parser) DateParser(dateFormat string) ParserType {
	return func(exp string) (time.Time, error) {
		date, err := time.Parse(dateFormat, exp)
		if err != nil {
			return time.Time{}, err
		}
		return date, nil
	}
}

func (p Parser) TimeParser(date time.Time, timeFormat string) ParserType {
	return func(exp string) (time.Time, error) {
		dtFormat := fmt.Sprintf("2006-01-02 %s", timeFormat)
		dtExp := fmt.Sprintf("%s %s", date.Format("2006-01-02"), exp)
		date, err := time.Parse(dtFormat, dtExp)
		if err != nil {
			return time.Time{}, err
		}
		return date, nil
	}
}

func (p Parser) ParseDuration(exp string) (time.Duration, error) {

	// default: hours
	if strings.HasSuffix(exp, "h") {

	} else if strings.HasSuffix(exp, "m") {

	} else {
		exp += "h"
	}

	// parse
	dur, err := time.ParseDuration(exp)
	if err != nil {
		return time.Duration(0), err
	}
	return dur, nil
}
