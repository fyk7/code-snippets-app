package model

import "time"

type User struct {
	UserID      uint64    `json:"user_id"`
	UserName    string    `json:"user_name"`
	IsSuperUser bool      `json:"is_super_user"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uint64    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   uint64    `json:"updated_by"`
}
