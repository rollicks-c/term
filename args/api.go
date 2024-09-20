package args

import (
	"fmt"
	"github.com/rollicks-c/term/internal/datetime"
	"github.com/rollicks-c/term/internal/num"
	"time"
)

const (
	DateInputFormat     = "060102"
	TimeInputFormat     = "1504"
	DateTimeInputFormat = "060102 1504"
)

type Batch struct {
	ac        *Collector
	errorList []error
}
type Collector struct {
	args           []string
	dateFormat     string
	timeFormat     string
	dateTimeFormat string
}

type ListProvider interface {
	SearchItems(exp string) ([]ListItem, error)
	ListItems() ([]ListItem, error)
}
type ListItem struct {
	Value any
	Name  string
}

type ArgContext[T any] struct {
	parsers      []argParser[T]
	value        T
	defaultValue *T
}

type CollectorOption func(*Collector)
type ArgOption[T any] func(ctx *ArgContext[T])

func WithDateLayout(layout string) CollectorOption {
	return func(c *Collector) {
		c.dateFormat = layout
	}
}

func WithDateTimeLayout(layout string) CollectorOption {
	return func(c *Collector) {
		c.dateTimeFormat = layout
	}
}

func WithDefault[T any](defaultValue T) ArgOption[T] {
	return func(ctx *ArgContext[T]) {
		ctx.defaultValue = &defaultValue
	}
}

func NewCollector(args []string, options ...CollectorOption) *Collector {
	ac := &Collector{
		args:           args,
		dateFormat:     DateInputFormat,
		timeFormat:     TimeInputFormat,
		dateTimeFormat: DateTimeInputFormat,
	}
	for _, opt := range options {
		opt(ac)
	}
	return ac
}

func (c Collector) Count() int {
	return len(c.args)
}

func (c Collector) Batch() *Batch {
	return &Batch{
		ac:        &c,
		errorList: make([]error, 0),
	}
}

func (b *Batch) Error() error {
	if len(b.errorList) == 0 {
		return nil
	}
	return fmt.Errorf("batch error: %v", b.errorList)
}

func (b *Batch) GetInt(index int, options ...ArgOption[int]) int {
	v, err := b.ac.GetInt(index, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (b *Batch) GetString(index int, options ...ArgOption[string]) string {
	v, err := b.ac.GetString(index, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (b *Batch) GetListItem(index int, provider ListProvider, options ...ArgOption[any]) any {
	v, err := b.ac.GetListItem(index, provider, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (b *Batch) GetDateAbs(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetDateAbs(index, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (b *Batch) GetDateRel(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetDateRel(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetDate(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetDate(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetWeek(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetWeek(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetWeekRel(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetWeekRel(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetMonth(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetMonth(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetMonthRel(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetMonthRel(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetDuration(index int, options ...ArgOption[time.Duration]) time.Duration {
	v, err := b.ac.GetDuration(index, options...)
	if err != nil {
		b.errorList = append(b.errorList)
	}
	return v
}

func (b *Batch) GetDateTimeAbs(index int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetDateTimeAbs(index, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (b *Batch) GetDateTimeRel(indexDate, indexTime int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetDateTimeRel(indexDate, indexTime, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (b *Batch) GetDateTime(indexDate, indexTime int, options ...ArgOption[time.Time]) time.Time {
	v, err := b.ac.GetDateTime(indexDate, indexTime, options...)
	if err != nil {
		b.errorList = append(b.errorList, err)
	}
	return v
}

func (c Collector) Validate(minCount int) error {
	if len(c.args) < minCount {
		return fmt.Errorf("at least %d args required", minCount)
	}
	return nil
}

func (c Collector) GetInt(index int, options ...ArgOption[int]) (int, error) {
	options = append(options, withParsers(num.ParseInt))
	return retrieve[int](c.args, index, options...)
}

func (c Collector) GetString(index int, options ...ArgOption[string]) (string, error) {
	options = append(options, withParsers(parseString))
	return retrieve[string](c.args, index, options...)
}

func (c Collector) GetListItem(index int, provider ListProvider, options ...ArgOption[any]) (any, error) {
	options = append(options, withParsers(itemSelector(provider)))
	return retrieve[any](c.args, index, options...)
}

func (c Collector) GetDate(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	dateRelParser := datetime.NewParser().ParseRelativeDay
	dateAbsParser := datetime.NewParser().DateParser(c.dateFormat)
	optionsDate := append(options, withParsers(dateRelParser, dateAbsParser))
	return retrieve[time.Time](c.args, index, optionsDate...)
}

func (c Collector) GetDateAbs(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	parser := datetime.NewParser().DateParser(c.dateFormat)
	options = append(options, withParsers(parser))
	return retrieve[time.Time](c.args, index, options...)
}

func (c Collector) GetDateTime(indexDate, indexTime int, options ...ArgOption[time.Time]) (time.Time, error) {

	// parse date part
	dateRelParser := datetime.NewParser().ParseRelativeDay
	dateAbsParser := datetime.NewParser().DateParser(c.dateFormat)
	optionsDate := append(options, withParsers(dateRelParser, dateAbsParser))
	date, err := retrieve[time.Time](c.args, indexDate, optionsDate...)
	if err != nil {
		return time.Time{}, err
	}

	// parse date with time
	timeParser := datetime.NewParser().TimeParser(date, c.timeFormat)
	optionsTime := append(options, withParsers(timeParser))
	return retrieve[time.Time](c.args, indexTime, optionsTime...)
}

func (c Collector) GetDateRel(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	parser := datetime.NewParser().ParseRelativeDay
	options = append(options, withParsers(parser))
	return retrieve[time.Time](c.args, index, options...)
}

func (c Collector) GetWeek(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	dateRelParser := datetime.NewParser().ParseRelativeWeek
	dateAbsParser := datetime.NewParser().DateParser(c.dateFormat)
	optionsDate := append(options, withParsers(dateRelParser, dateAbsParser))
	return retrieve[time.Time](c.args, index, optionsDate...)
}

func (c Collector) GetWeekRel(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	parser := datetime.NewParser().ParseRelativeWeek
	options = append(options, withParsers(parser))
	return retrieve[time.Time](c.args, index, options...)
}

func (c Collector) GetMonth(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	dateRelParser := datetime.NewParser().ParseRelativeMonth
	dateAbsParser := datetime.NewParser().DateParser(c.dateFormat)
	optionsDate := append(options, withParsers(dateRelParser, dateAbsParser))
	return retrieve[time.Time](c.args, index, optionsDate...)
}

func (c Collector) GetMonthRel(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	parser := datetime.NewParser().ParseRelativeMonth
	options = append(options, withParsers(parser))
	return retrieve[time.Time](c.args, index, options...)
}

func (c Collector) GetDateTimeAbs(index int, options ...ArgOption[time.Time]) (time.Time, error) {
	parser := datetime.NewParser().DateParser(c.dateTimeFormat)
	options = append(options, withParsers(parser))
	return retrieve[time.Time](c.args, index, options...)
}

func (c Collector) GetDateTimeRel(indexDate, indexTime int, options ...ArgOption[time.Time]) (time.Time, error) {

	// parse date part
	dateParser := datetime.NewParser().ParseRelativeDay
	optionsDate := append(options, withParsers(dateParser))
	date, err := retrieve[time.Time](c.args, indexDate, optionsDate...)
	if err != nil {
		return time.Time{}, err
	}

	// parse date with time
	timeParser := datetime.NewParser().TimeParser(date, c.timeFormat)
	optionsTime := append(options, withParsers(timeParser))
	return retrieve[time.Time](c.args, indexTime, optionsTime...)
}

func (c Collector) GetDuration(index int, options ...ArgOption[time.Duration]) (time.Duration, error) {
	parser := datetime.NewParser().ParseDuration
	options = append(options, withParsers(parser))
	return retrieve[time.Duration](c.args, index, options...)
}
