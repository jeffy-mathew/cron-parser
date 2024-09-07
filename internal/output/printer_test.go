package output

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/jeffy-mathew/cron-parser/parser"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update golden files")

func TestPrintSchedule(t *testing.T) {
	schedule := parser.Schedule{
		MinutesField: parser.Field{
			Values: []int{0, 15, 30, 45},
		},
		HoursField: parser.Field{
			Values: []int{0},
		},
		DaysOfMonthField: parser.Field{
			Values: []int{1, 15},
		},
		MonthsField: parser.Field{
			Values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		DaysOfWeekField: parser.Field{
			Values: []int{1, 2, 3, 4, 5},
		},
		Command: `/usr/bin/echo "Hello, World!"`,
	}

	result := PrintSchedule(schedule)

	goldenFile := filepath.Join("testdata", "print_schedule.golden")
	if *update {
		err := os.WriteFile(goldenFile, []byte(result), 0644)
		assert.NoError(t, err)
	}

	expected, err := os.ReadFile(goldenFile)
	assert.NoError(t, err)

	assert.Equal(t, string(expected), PrintSchedule(schedule))

}
