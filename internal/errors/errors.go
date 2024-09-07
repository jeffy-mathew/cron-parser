package errors

import (
	"errors"
	"fmt"
)

var (
	New = errors.New
)

var (
	ErrInvalidCronExpression = errors.New("invalid cron expression")
	ErrInvalidParserType     = errors.New("invalid parser type")
)

type ErrOutOfBounds struct {
	Value int
	Lower int
	Upper int
}

func (e ErrOutOfBounds) Error() string {
	return fmt.Sprintf("value %d is out of bounds [%d, %d]", e.Value, e.Lower, e.Upper)
}
