package main

import "time"

type ToDo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateTime    time.Time `json:"dateTime"`
	Completed   bool      `json:"completed"`
}
