package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jeffy-mathew/cron-parser/internal/output"
	"github.com/jeffy-mathew/cron-parser/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a cron expression as an argument.")
		os.Exit(1)
	}

	cronExpression := strings.Join(os.Args[1:], " ")

	p, err := parser.NewParser(parser.StdParser)
	if err != nil {
		fmt.Printf("Error creating parser: %v\n", err)
		os.Exit(1)
	}

	schedule, err := p.Parse(cronExpression)
	if err != nil {
		fmt.Printf("Error parsing cron expression: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(output.PrintSchedule(schedule))
}
