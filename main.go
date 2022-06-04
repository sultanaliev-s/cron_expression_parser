package main

import (
	"fmt"
	"os"

	"github.com/sultanaliev-s/cron_expression_parser/cronparser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(
			"Invalid number of arguments. Expected 1. Got", len(os.Args)-1)
		return
	}

	p, err := cronparser.New(os.Args[1])
	if err != nil {
		fmt.Println("Invalid cron expression")
		return
	}

	fmt.Println(p)
}
