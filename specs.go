package main

import "time"

//go:generate reform

//reform:specs
type Spec struct {
	Status     int64     `reform:"status"`
	Filename   string    `reform:"filename"`
	LineNumber int64     `reform:"line_number"`
	CommitID   int64     `reform:"commit_id"`
	CreatedAt  time.Time `reform:"created_at"`
}
