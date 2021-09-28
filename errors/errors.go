package main

import (
	"fmt"
	"time"
)

type MyError struct {
	Message   string
	Timestamp time.Time
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v %s", e.Message, e.Timestamp)
}

func run() error {
	return &MyError{
		Message:   "Unhandled problem",
		Timestamp: time.Now(),
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
