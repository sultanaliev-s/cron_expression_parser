# Cron Expression Parser
An application which parses a cron string and expands each field to show the
times at which it will run.

Supports five time fields (minute, hour, day of month, month, and day of week)
plus a command. Does not support multiple patterns per field (e.g. "1,5-10,20"),
special time strings such as "@yearly".

The input should be as the command line argument consisting of 1 string.

## Getting started

1. Install Go as described [here](https://go.dev/doc/install)
2. Go to the project root directory and run 
```
go build main.go
``` 
3. Run the application
```
./main "* * * * * <your command>"
```
