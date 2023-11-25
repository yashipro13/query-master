package repository

import "time"

type project struct {
	id          int
	name        string
	slug        string
	description string
	createdAt   time.Time
}

