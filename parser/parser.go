package parser

import (
	"github.com/jeffy-mathew/cron-parser/internal/errors"
)

const (
	// StdParser is the standard parser for cron expressions.
	StdParser = iota
)

// CronScheduleParser is an interface for parsing cron expressions.
type CronScheduleParser interface {
	Parse(expression string) (Schedule, error)
}

// NewParser returns a new CronScheduleParser.
func NewParser(parserType int) (CronScheduleParser, error) {
	switch parserType {
	case StdParser:
		return NewStandardCronParser(), nil
	default:
		return nil, errors.ErrInvalidParserType
	}
}
