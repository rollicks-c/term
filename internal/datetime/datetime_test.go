package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateOnly(t *testing.T) {

	now := time.Now()
	act := DateOnly(now)
	assert.Equal(t, now.Year(), act.Year())
	assert.Equal(t, now.Month(), act.Month())
	assert.Equal(t, now.Day(), act.Day())
	assert.Equal(t, 0, act.Hour())
	assert.Equal(t, 0, act.Minute())
	assert.Equal(t, 0, act.Second())
}

func TestRelDay(t *testing.T) {
	now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	p := NewParser(WithNow(now))
	{
		act, err := p.ParseRelativeDay("-")
		assert.NoError(t, err)
		assert.Equal(t, 2021, act.Year())
		assert.Equal(t, 1, int(act.Month()))
		assert.Equal(t, 1, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
	{
		act, err := p.ParseRelativeDay("2-")
		assert.NoError(t, err)
		assert.Equal(t, 2020, act.Year())
		assert.Equal(t, 12, int(act.Month()))
		assert.Equal(t, 30, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
	{
		act, err := p.ParseRelativeDay("4+")
		assert.NoError(t, err)
		assert.Equal(t, 2021, act.Year())
		assert.Equal(t, 1, int(act.Month()))
		assert.Equal(t, 5, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
}

func TestRelWeek(t *testing.T) {
	now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	p := NewParser(WithNow(now))
	{
		act, err := p.ParseRelativeWeek("-")
		assert.NoError(t, err)
		assert.Equal(t, 2021, act.Year())
		assert.Equal(t, 1, int(act.Month()))
		assert.Equal(t, 1, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
	{
		act, err := p.ParseRelativeWeek("2-")
		assert.NoError(t, err)
		assert.Equal(t, 2020, act.Year())
		assert.Equal(t, 12, int(act.Month()))
		assert.Equal(t, 18, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
	{
		act, err := p.ParseRelativeWeek("4+")
		assert.NoError(t, err)
		assert.Equal(t, 2021, act.Year())
		assert.Equal(t, 1, int(act.Month()))
		assert.Equal(t, 29, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
}

func TestDateParse(t *testing.T) {
	{
		p := NewParser().DateParser("2006-01-02")
		act, err := p("2021-01-01")
		assert.NoError(t, err)
		assert.Equal(t, 2021, act.Year())
		assert.Equal(t, 1, int(act.Month()))
		assert.Equal(t, 1, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
	{
		p := NewParser().DateParser("060102")
		act, err := p("210101")
		assert.NoError(t, err)
		assert.Equal(t, 2021, act.Year())
		assert.Equal(t, 1, int(act.Month()))
		assert.Equal(t, 1, act.Day())
		assert.Equal(t, 0, act.Hour())
		assert.Equal(t, 0, act.Minute())
		assert.Equal(t, 0, act.Second())
	}
}

func TestDurationParse(t *testing.T) {

	p := NewParser()
	{
		act, err := p.ParseDuration("3m")
		assert.NoError(t, err)
		assert.Equal(t, time.Minute*3, act)
	}
	{
		act, err := p.ParseDuration("5h")
		assert.NoError(t, err)
		assert.Equal(t, time.Hour*5, act)
	}
	{
		act, err := p.ParseDuration("7.75h")
		assert.NoError(t, err)
		assert.Equal(t, time.Hour*7+time.Minute*45, act)
	}
	{
		act, err := p.ParseDuration("12.25")
		assert.NoError(t, err)
		assert.Equal(t, time.Hour*12+time.Minute*15, act)
	}
}
