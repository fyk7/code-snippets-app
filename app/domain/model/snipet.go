package model

import "time"

type Snippet struct {
	SnippetID          uint64 `json:"snipet_id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Body               string `json:"body"`
	ProgramingLanguage string `json:"programing_language"`
	// AuthorID           uint64    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uint64    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uint64    `json:"updated_by"`
}
