package parser_test

import (
	"testing"

	"github.com/jeffy-mathew/cron-parser/internal/errors"
	"github.com/jeffy-mathew/cron-parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestCronScheduleParser_Parse(t *testing.T) {
	tests := []struct {
		name           string
		cronExpression string
		wantMinutes    func() []int
		wantHours      func() []int
		wantDayOfMonth func() []int
		wantMonths     func() []int
		wantDayOfWeek  func() []int
		wantCommand    string
		wantErr        error
	}{
		{
			name:           "invalid cron expression",
			cronExpression: "invalid",
			wantErr:        errors.ErrInvalidCronExpression,
		},
		{
			name:           "invalid cron expression in minutes",
			cronExpression: "m * * * * /usr/bin/time",
			wantErr:        errors.ErrInvalidCronExpression,
		},
		{
			name:           "invalid cron expression in every hour",
			cronExpression: "* */s * * * /usr/bin/time",
			wantErr:        errors.ErrInvalidCronExpression,
		},
		{
			name:           "specifying range - out of bound",
			cronExpression: "* */2 0-3 * * /usr/bin/time",
			wantErr:        errors.ErrOutOfBounds{Value: 0, Lower: 1, Upper: 31},
		},
		{
			name:           "invalid cron expression - above upper bound",
			cronExpression: "60 */2 1-3 * * /usr/bin/time",
			wantErr:        errors.ErrOutOfBounds{Value: 60, Lower: 0, Upper: 59},
		},
		{
			name:           "invalid cron expression - below lower bound",
			cronExpression: "59 */2 1-3 0 * /bin/bash/time",
			wantErr:        errors.ErrOutOfBounds{Value: 0, Lower: 1, Upper: 12},
		},
		{
			name:           "invalid cron expression while specifying step",
			cronExpression: "59 */2 1-3 3/2 * /usr/bin/time",
			wantErr:        errors.ErrInvalidCronExpression,
		},
		{
			name:           "invalid cron expression while specifying range",
			cronExpression: "59 */2 1-3-4 3 * /usr/bin/time",
			wantErr:        errors.ErrInvalidCronExpression,
		},
		{
			name:           "invalid value in range",
			cronExpression: "59 */2 10 3 a-b /usr/bin/time",
			wantErr:        errors.ErrInvalidCronExpression,
		},
		{
			name:           "simple cron expression",
			cronExpression: "0 15 10 * * /usr/bin/echo hello",
			wantMinutes:    func() []int { return []int{0} },
			wantHours:      func() []int { return []int{15} },
			wantDayOfMonth: func() []int { return []int{10} },
			wantMonths: func() []int {
				months := make([]int, 12)
				for i := 1; i <= 12; i++ {
					months[i-1] = i
				}

				return months
			},
			wantDayOfWeek: func() []int {
				weeks := make([]int, 7)
				for i := 0; i < 7; i++ {
					weeks[i] = i
				}

				return weeks
			},
			wantCommand: "/usr/bin/echo hello",
		},
		{
			name:           "Every minute",
			cronExpression: "* * * * * /usr/bin/time",
			wantMinutes: func() []int {
				minutes := make([]int, 60)
				for i := 0; i < 60; i++ {
					minutes[i] = i
				}

				return minutes
			},
			wantHours: func() []int {
				hours := make([]int, 24)
				for i := 0; i < 24; i++ {
					hours[i] = i
				}

				return hours
			},
			wantDayOfMonth: func() []int {
				days := make([]int, 31)
				for i := 1; i <= 31; i++ {
					days[i-1] = i
				}

				return days
			},
			wantMonths: func() []int {
				months := make([]int, 12)
				for i := 1; i <= 12; i++ {
					months[i-1] = i
				}

				return months
			},
			wantDayOfWeek: func() []int {
				weeks := make([]int, 7)
				for i := 0; i < 7; i++ {
					weeks[i] = i
				}

				return weeks
			},
			wantCommand: "/usr/bin/time",
		},
		{
			name:           "every 2 hours",
			cronExpression: "0 */2 * * * /usr/bin/time",
			wantMinutes:    func() []int { return []int{0} },
			wantHours: func() []int {
				hours := make([]int, 12)
				for i := 0; i < 24; i++ {
					if i%2 == 0 {
						hours[i/2] = i
					}
				}

				return hours
			},
			wantDayOfMonth: func() []int {
				days := make([]int, 31)
				for i := 1; i <= 31; i++ {
					days[i-1] = i
				}

				return days
			},
			wantMonths: func() []int {
				months := make([]int, 12)
				for i := 1; i <= 12; i++ {
					months[i-1] = i
				}

				return months
			},
			wantDayOfWeek: func() []int {
				weeks := make([]int, 7)
				for i := 0; i < 7; i++ {
					weeks[i] = i
				}

				return weeks
			},
			wantCommand: "/usr/bin/time",
		},
	}

	scheduleParser, err := parser.NewParser(parser.StdParser)
	assert.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scheduleParser.Parse(tt.cronExpression)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			assert.EqualValues(t, tt.wantMinutes(), got.Minutes())
			assert.EqualValues(t, tt.wantHours(), got.Hours())
			assert.EqualValues(t, tt.wantDayOfMonth(), got.DaysOfMonth())
			assert.EqualValues(t, tt.wantMonths(), got.Months())
			assert.EqualValues(t, tt.wantDayOfWeek(), got.DaysOfWeek())
			assert.Equal(t, tt.wantCommand, got.Command)
		})
	}
}
