package output

import (
	"fmt"
	"strings"

	"github.com/jeffy-mathew/cron-parser/parser"
)

// PrintSchedule formats and prints the schedule
func PrintSchedule(s parser.Schedule) string {
	var output strings.Builder
	fieldNames := []string{"minute", "hour", "day of month", "month", "day of week", "command"}
	fieldValues := [][]int{
		s.Minutes(),
		s.Hours(),
		s.DaysOfMonth(),
		s.Months(),
		s.DaysOfWeek(),
	}

	for i, name := range fieldNames {
		if i < 5 {
			fmt.Fprintf(&output, "%-14s", name)
			for _, v := range fieldValues[i] {
				fmt.Fprintf(&output, "%d ", v)
			}
			fmt.Fprintln(&output)
		} else {
			fmt.Fprintf(&output, "%-14s%s\n", name, s.Cmd())
		}
	}

	return output.String()
}
