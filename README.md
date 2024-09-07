# Cron Expression Parser

This project is a command-line application that parses cron expressions and displays the schedule in a human-readable format.

## Features

- Parses standard cron expressions
- Supports the five time fields (minute, hour, day of month, month, day of week) and a command
- Displays the schedule in a formatted table
- Handles various cron expression formats, including ranges, lists, and step values

## Installation

To install the cron parser, make sure you have Go installed on your system.
Alternatively you can use [Docker container](#docker-usage) to run the application.

To run the application with Go installed, run:

```
go install github.com/jeffy-mathew/cron-parser/cmd/cron-parser@latest
```

Or clone the repository and run (assumes $GOPATH/bin is in your $PATH):

```
go build -o cron-parser ./cmd/main.go
cp cron-parser $GOPATH/bin/
```

## Usage

After installation, you can run the parser using the following command:


From source build/go install:
```
cron-parser "*/15 0 1,15 * 1-5 /usr/bin/echo hello"
```

Replace the cron expression in quotes with your own expression.


## Docker Usage

You can also run the cron parser using [Docker Container](https://www.docker.com/). First, build the Docker image:

```
docker build -t cron-parser .
```

Then, run the parser using Docker:

```
docker run cron-parser "*/15 0 1,15 * 1-5 /usr/bin/echo hello"
```

## Output

The parser will output a formatted table showing the schedule for each field.
For the input provided above the output will be:

```
minute        0 15 30 45 
hour          0 
day of month  1 15 
month         1 2 3 4 5 6 7 8 9 10 11 12 
day of week   1 2 3 4 5 
command       /usr/bin/echo hello
```


## Development

### Project Structure

- `cmd/main.go`: Entry point of the application
- `parser/`: Contains the cron expression parsing logic, this can be used as a library in other projects.
- `internal/output/`: Handles the formatting and printing of the schedule
- `internal/output/testdata/`: Contains golden files for output testing

### Running Tests

To run the tests for this project:

```
go test ./...
```

To update the golden files used in testing:

```
go test ./internal/output -update
```

## Taskfile usage

This project also uses [Taskfile](https://taskfile.dev/) to manage build and run commands. Make sure you have Taskfile installed on your system to make use of it, alternatively you can use the commands directly used in the Taskfile or follow the instructions described above to build and run the application from source.

To run the cron parser from source build, using Taskfile:

```
task run -- "*/15 0 1,15 * 1-5 /usr/bin/echo hello"
```

To run the cron parser from docker, using Taskfile:

```
task docker -- "*/15 0 1,15 * 1-5 /usr/bin/echo hello"
```
