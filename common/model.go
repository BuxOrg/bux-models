package common

import "time"

type Model struct {
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	DeletedAt time.Time              `json:"deleted_at"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// TODO: add rawXpub to xpub model
