package parser_test

import (
	"testing"

	"github.com/jeffy-mathew/cron-parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestNewCronSchedule(t *testing.T) {
	schedule := parser.NewSchedule()
	assert.NotEmpty(t, schedule)

	tests := []struct {
		name     string
		field    parser.Field
		expected parser.Bound
	}{
		{"Minutes", schedule.MinutesField, parser.MinutesBound},
		{"Hours", schedule.HoursField, parser.HoursBound},
		{"DaysOfMonth", schedule.DaysOfMonthField, parser.DayOfMonthBound},
		{"Months", schedule.MonthsField, parser.MonthsBound},
		{"DaysOfWeek", schedule.DaysOfWeekField, parser.DayOfWeekBound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.field.Bound)
		})
	}
}

func TestCronField_AddValue(t *testing.T) {
	tests := []struct {
		name    string
		field   parser.Field
		value   int
		wantErr bool
	}{
		{"Valid value", parser.Field{Bound: parser.MinutesBound}, 30, false},
		{"Lower bound", parser.Field{Bound: parser.MinutesBound}, 0, false},
		{"Upper bound", parser.Field{Bound: parser.MinutesBound}, 59, false},
		{"Below lower bound", parser.Field{Bound: parser.MinutesBound}, -1, true},
		{"Above upper bound", parser.Field{Bound: parser.MinutesBound}, 60, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.field.AddValue(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Contains(t, tt.field.Vals(), tt.value)
		})
	}
}


func TestCronField_SetValues(t *testing.T) {
	tests := []struct {
		name    string
		field   parser.Field
		values  []int
		wantErr bool
	}{
		{"Valid values", parser.Field{Bound: parser.HoursBound}, []int{0, 12, 23}, false},
		{"Invalid value", parser.Field{Bound: parser.HoursBound}, []int{0, 24}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.field.SetValues(tt.values)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tt.values, tt.field.Values)
		})
	}
}