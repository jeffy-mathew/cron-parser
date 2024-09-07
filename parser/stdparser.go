package parser

import (
	"strconv"
	"strings"

	"github.com/jeffy-mathew/cron-parser/internal/errors"
)

// StandardCronParser is a parser for cron expressions.
type StandardCronParser struct {
}

// NewStandardCronParser returns a new StandardCronParser.
func NewStandardCronParser() CronScheduleParser {
	return &StandardCronParser{}
}

// Parse parses a cron expression and returns a Schedule.
func (p *StandardCronParser) Parse(expression string) (Schedule, error) {
	fields := strings.Fields(expression)
	if len(fields) < 6 {
		return Schedule{}, errors.ErrInvalidCronExpression
	}

	schedule := NewSchedule()
	err := p.parseField(fields[0], &schedule.MinutesField)
	if err != nil {
		return Schedule{}, err
	}

	err = p.parseField(fields[1], &schedule.HoursField)
	if err != nil {
		return Schedule{}, err
	}

	err = p.parseField(fields[2], &schedule.DaysOfMonthField)
	if err != nil {
		return Schedule{}, err
	}

	err = p.parseField(fields[3], &schedule.MonthsField)
	if err != nil {
		return Schedule{}, err
	}

	err = p.parseField(fields[4], &schedule.DaysOfWeekField)
	if err != nil {
		return Schedule{}, err
	}

	schedule.Command = strings.Join(fields[5:], " ")
	return schedule, nil
}

// parseField parses a single field of the cron expression
func (p *StandardCronParser) parseField(field string, cronField *Field) error {
	if field == "*" {
		for i := int(cronField.LowerBound); i <= int(cronField.UpperBound); i++ {
			if err := cronField.AddValue(i); err != nil {
				return err
			}
		}
		return nil
	}

	parts := strings.Split(field, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return errors.ErrInvalidCronExpression
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return errors.ErrInvalidCronExpression
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return errors.ErrInvalidCronExpression
			}

			for i := start; i <= end; i++ {
				if err := cronField.AddValue(i); err != nil {
					return err
				}
			}
		} else if strings.Contains(part, "/") {
			stepParts := strings.Split(part, "/")
			if len(stepParts) != 2 || stepParts[0] != "*" {
				return errors.ErrInvalidCronExpression
			}

			step, err := strconv.Atoi(stepParts[1])
			if err != nil {
				return errors.ErrInvalidCronExpression
			}

			for i := int(cronField.LowerBound); i <= int(cronField.UpperBound); i += step {
				if err := cronField.AddValue(i); err != nil {
					return err
				}
			}
		} else {
			value, err := strconv.Atoi(part)
			if err != nil {
				return errors.ErrInvalidCronExpression
			}

			if err := cronField.AddValue(value); err != nil {
				return err
			}
		}
	}

	return nil
}
