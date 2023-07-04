package buxmodels

import "time"

type SyncResults struct {
	LastMessage string        `json:"last_message"`
	Results     []*SyncResult `json:"results"`
}

type SyncResult struct {
	Action        string    `json:"action"`             // type: broadcast, sync etc
	ExecutedAt    time.Time `json:"executed_at"`        // Time it was executed
	Provider      string    `json:"provider,omitempty"` // Provider used for attempt(s)
	StatusMessage string    `json:"status_message"`     // Success or failure message
}
