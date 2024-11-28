package data

import (
	"time"
)

type BlogPost struct {
	ID 			int64
	CreatedAt 	time.Time
	Title       string
	Description string
	Content     string
	Tags        []string
	Slug        string
	Version 	int32
}