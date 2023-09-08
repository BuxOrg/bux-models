// Package buxmodels contains all models (contracts) between bux-server api and other bux solutions
package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

// AccessKey is a model that represents an access key.
type AccessKey struct {
	// Model is a common model that contains common fields for all models.
	Model common.Model

	// ID is an access key id.
	ID string `json:"id"`
	// XpubID is an access key's xpub related id.
	XpubID string `json:"xpub_id"`
	// RevokedAt is a time when access key was revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// Key is a string representation of an access key.
	Key string `json:"key,omitempty"`
}
