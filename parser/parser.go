package parser

// CronScheduleParser is an interface for parsing cron expressions.
type CronScheduleParser interface {
	Parse(expression string) (Schedule, error)
}
