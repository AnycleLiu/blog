package core

import "time"

type Article struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Tag         string    `json:"tag"`
	CreatedTime time.Time `json:"createdTime" binding:"required"`
}
