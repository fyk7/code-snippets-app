package model

import "time"

type Tag struct {
	TagID     uint64    `json:"tag_id"`
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uint64    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uint64    `json:"updated_by"`
}
