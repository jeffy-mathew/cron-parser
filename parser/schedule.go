package parser

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

// Schedule is a struct that holds the cron schedule.
type Schedule struct {
	MinutesField     Field
	HoursField       Field
	DaysOfMonthField Field
	MonthsField      Field
	DaysOfWeekField  Field
	Command          string
}
