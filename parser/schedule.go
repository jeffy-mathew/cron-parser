package parser

import (
	"github.com/jeffy-mathew/cron-parser/internal/errors"
)

var (
	MinutesBound    = Bound{LowerBound: 0, UpperBound: 59}
	HoursBound      = Bound{LowerBound: 0, UpperBound: 23}
	DayOfMonthBound = Bound{LowerBound: 1, UpperBound: 31}
	MonthsBound     = Bound{LowerBound: 1, UpperBound: 12}
	DayOfWeekBound  = Bound{LowerBound: 0, UpperBound: 6}
)

// Bound is a struct that holds the lower and upper bounds.
type Bound struct {
	LowerBound int
	UpperBound int
}

// Field is a generic struct to hold values and bounds.
type Field struct {
	Values []int
	Bound
}

// Vals returns the values of the cronField.
func (c Field) Vals() []int {
	return c.Values
}

// AddValue adds a value to the cronField, enforcing bounds
func (c *Field) AddValue(value int) error {
	if value < int(c.LowerBound) || value > int(c.UpperBound) {
		return errors.ErrOutOfBounds{Value: value, Lower: int(c.LowerBound), Upper: int(c.UpperBound)}
	}

	c.Values = append(c.Values, value)
	return nil
}

// SetValues sets multiple values, enforcing bounds on all
func (c *Field) SetValues(values []int) error {
	for _, v := range values {
		if err := c.AddValue(v); err != nil {
			return err
		}
	}

	return nil
}


// Schedule is a struct that holds the cron schedule.
type Schedule struct {
	MinutesField     Field
	HoursField       Field
	DaysOfMonthField Field
	MonthsField      Field
	DaysOfWeekField  Field
	Command          string
}

// NewSchedule creates a new Schedule with all boundaries initialized.
func NewSchedule() Schedule {
	return Schedule{
		MinutesField:     Field{Bound: MinutesBound},
		HoursField:       Field{Bound: HoursBound},
		DaysOfMonthField: Field{Bound: DayOfMonthBound},
		MonthsField:      Field{Bound: MonthsBound},
		DaysOfWeekField:  Field{Bound: DayOfWeekBound},
	}
}
