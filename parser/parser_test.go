package parser

import (
	"testing"

	"github.com/jeffy-mathew/cron-parser/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	tests := []struct {
		name       string
		parserType int
		wantType   CronScheduleParser
		wantErr    error
	}{
		{
			name:       "Standard Parser",
			parserType: StdParser,
			wantType:   &StandardCronParser{},
			wantErr:    nil,
		},
		{
			name:       "Invalid Parser Type",
			parserType: 999, // An invalid parser type
			wantType:   nil,
			wantErr:    errors.ErrInvalidParserType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewParser(tt.parserType)

			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}

			assert.IsType(t, tt.wantType, got)
		})
	}
}
