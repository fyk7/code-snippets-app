package model

import "time"

type SnippetTagRelation struct {
	SnippetTagRelationID uint64    `json:"snippet_tag_relation_id"`
	SnippetID            uint64    `json:"snippet_id"`
	TagID                uint64    `json:"tag_id"`
	CreatedAt            time.Time `json:"created_at"`
	CreatedBy            uint64    `json:"created_by"`
	UpdatedAt            time.Time `json:"updated_at"`
	UpdatedBy            uint64    `json:"updated_by"`
}
