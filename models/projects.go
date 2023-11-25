package models

import "time"

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	Hashtags    []string  `json:"hashtags"`
	CreatedAt   time.Time `json:"createdAt"`
}
